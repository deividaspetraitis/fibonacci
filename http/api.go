package http

import (
	"context"
	"net/http"
	stdhttp "net/http"
	"os"
	"syscall"

	"github.com/deividaspetraitis/fibonacci"
	"github.com/deividaspetraitis/fibonacci/log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct
type App struct {
	API      *mux.Router
	shutdown chan os.Signal
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal) *App {
	api := App{
		API:      mux.NewRouter(),
		shutdown: shutdown,
	}
	return &api
}

// ServeHTTP API
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.API.ServeHTTP(w, r)
}

// SignalShutdown is used to gracefully shutdown the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

// API constructs an http.Handler with all application routes defined.
func API(shutdown chan os.Signal, cfg *Config, app *fibonacci.Fibonacci, logger log.Logger) stdhttp.Handler {
	// =========================================================================
	// Construct the web app api which holds all routes as well as common Middleware.

	api := NewApp(shutdown)

	// =========================================================================
	// Construct and attach relevant handlers to web app api

	api.API.HandleFunc("/current", GetCurrentFibonacciNumber(func(ctx context.Context) (int64, error) {
		return app.CurrentFibonacciNumber(ctx)
	})).Methods(http.MethodGet)

	api.API.HandleFunc("/next", GetNextFibonacciNumberFunc(func(ctx context.Context) (int64, error) {
		return app.NextFibonacciNumber(ctx)
	})).Methods(http.MethodGet)

	api.API.HandleFunc("/previous", GetPreviousFibonacciNumberFunc(func(ctx context.Context) (int64, error) {
		return app.PreviousFibonacciNumber(ctx)
	})).Methods(http.MethodGet)

	router := mux.NewRouter()

	// recover from a panic, log, and continue to the next handler
	router.PathPrefix("/").Handler(handlers.RecoveryHandler()(api.API))

	return router
}
