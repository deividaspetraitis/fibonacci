package http

import (
	"context"
	"net/http"

	"github.com/deividaspetraitis/fibonacci"
	"github.com/deividaspetraitis/fibonacci/errors"
	"github.com/deividaspetraitis/fibonacci/log"
	"github.com/deividaspetraitis/fibonacci/pkg/api/v1"
)

// getFibonacciNumberFunc decouples actual Fibonacci number retrieval implementation and allows easily test HTTP handler.
type getFibonacciNumberFunc func(ctx context.Context) (int64, error)

// GetCurrentFibonacciNumberFunc responds with the current number in the Fibonacci sequence.
func GetCurrentFibonacciNumber(getCurrentFibonacciNumber getFibonacciNumberFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// It's always json.
		w.Header().Set("Content-Type", "application/json")

		number, err := getCurrentFibonacciNumber(r.Context())
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"handler": "fibonacci",
				"method":  "GetCurrentFibonacciNumber",
			}).Println("encountered an error retrieving Fibonacci number")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := api.CurrentFibonacciNumberResponse{
			Current: number,
		}

		w.WriteHeader(http.StatusOK)
		if err := Marshal(w, &response); err != nil {
			log.WithError(err).WithFields(log.Fields{
				"handler": "fibonacci",
				"method":  "GetCurrentFibonacciNumber",
			}).Println("unable to marshal response data")

			return
		}
	}
}

// GetNextFibonacciNumberFunc responds with the next number in the Fibonacci sequence.
func GetNextFibonacciNumberFunc(getNextFibonacciNumber getFibonacciNumberFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// It's always json.
		w.Header().Set("Content-Type", "application/json")

		number, err := getNextFibonacciNumber(r.Context())
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"handler": "fibonacci",
				"method":  "GetNextFibonacciNumberFunc",
			}).Println("encountered an error retrieving Fibonacci number")

			w.WriteHeader(http.StatusInternalServerError)

			if errors.Is(err, fibonacci.ErrCounterOverflow) {
				Marshal(w, &api.Error{ 
					Message: "counter overflow",
				})
			}
			return
		}

		response := api.NextFibonacciNumberResponse{
			Next: number,
		}

		w.WriteHeader(http.StatusOK)
		if err := Marshal(w, &response); err != nil {
			log.WithError(err).WithFields(log.Fields{
				"handler": "fibonacci",
				"method":  "GetNextFibonacciNumberFunc",
			}).Println("unable to marshal response data")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// GetPreviousFibonacciNumberFunc responds with the next number in the Fibonacci sequence.
func GetPreviousFibonacciNumberFunc(getPreviousFibonacciNumber getFibonacciNumberFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// It's always json.
		w.Header().Set("Content-Type", "application/json")

		number, err := getPreviousFibonacciNumber(r.Context())
		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"handler": "fibonacci",
				"method":  "GetPreviousFibonacciNumberFunc",
			}).Println("encountered an error retrieving Fibonacci number")

			w.WriteHeader(http.StatusInternalServerError)

			if errors.Is(err, fibonacci.ErrCounterUnderflow) {
				Marshal(w, &api.Error{
					Message: "counter underflow",
				})
			}
			return
		}

		response := api.PreviousFibonacciNumberResponse{
			Previous: number,
		}

		w.WriteHeader(http.StatusOK)
		if err := Marshal(w, &response); err != nil {
			log.WithError(err).WithFields(log.Fields{
				"handler": "fibonacci",
				"method":  "GetPreviousFibonacciNumberFunc",
			}).Println("unable to marshal response data")

			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
