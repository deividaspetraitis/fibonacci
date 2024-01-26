package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/deividaspetraitis/fibonacci"
	"github.com/deividaspetraitis/fibonacci/config"
	"github.com/deividaspetraitis/fibonacci/errors"
	ihttp "github.com/deividaspetraitis/fibonacci/http"
	"github.com/deividaspetraitis/fibonacci/log"
)

var shutdowntimeout = time.Duration(5) * time.Second

// program flags
var (
	cfgPath string
)

// initialise program state
func init() {
	flag.StringVar(&cfgPath, "config", os.Getenv("config"), "PATH to .env configuration file")
}

// main program entry point.
func main() {
	flag.Parse()

	logger := log.Default()

	cfg, err := config.New(cfgPath)
	if err != nil {
		logger.WithError(err).Fatal("parsing configuration file")
	}

	if err := run(cfg, logger); err != nil {
		logger.WithError(err).Fatal("unable to start service")
	}
}

func run(cfg *config.Config, logger log.Logger) error {
	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// =========================================================================
	// Construct services

	app := fibonacci.Fibonacci{}

	// =========================================================================
	// Start HTTP server

	api := http.Server{
		Addr:    cfg.HTTP.Address,
		Handler: ihttp.API(shutdown, cfg.HTTP, &app, logger),
	}

	go func() {
		logger.Printf("http server listening on %s", cfg.HTTP.Address)
		serverErrors <- api.ListenAndServe()
	}()

	// ========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		logger.Printf("http server start shutdown caused by %v", sig)

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), shutdowntimeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := api.Shutdown(ctx)
		if err != nil {
			logger.WithError(err).Error("graceful shutdown did not complete")
			api.Close()
		}

		// Log the status of this shutdown.
		switch {
		case sig == syscall.SIGSTOP:
			return errors.New("integrity issue caused shutdown")
		case err != nil:
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
