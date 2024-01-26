package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/deividaspetraitis/fibonacci"
	"github.com/deividaspetraitis/fibonacci/errors"

	"github.com/gorilla/mux"
)

func TestGetCurrentFibonacciNumber(t *testing.T) {
	var testcases = []struct {
		getFibonacciNumber getFibonacciNumberFunc

		response   string
		statusCode int
	}{
		// result response
		{
			getFibonacciNumber: func(ctx context.Context) (int64, error) {
				return 1, nil
			},
			response:   `{"current":1}`,
			statusCode: http.StatusOK,
		},
		// service error
		{
			getFibonacciNumber: func(ctx context.Context) (int64, error) {
				return 0, errors.New("test getFibonacciNumber error")
			},
			response:   "",
			statusCode: http.StatusInternalServerError,
		},
	}

	for i, tt := range testcases {
		req := httptest.NewRequest(http.MethodPost, "http://localhost/current", nil)
		w := httptest.NewRecorder()

		// To add the vars to the context we need to create a router through which we can pass the request.
		// TODO: tests should be not aware of routing mechanism?
		router := mux.NewRouter()
		router.HandleFunc("/current", GetCurrentFibonacciNumber(tt.getFibonacciNumber))

		router.ServeHTTP(w, req)

		if statusCode := w.Result().StatusCode; statusCode != tt.statusCode {
			t.Errorf("#%d HTTP status got %v, want %v", i, statusCode, tt.statusCode)
		}

		// we do apply TrimSpace to clean up response coming from HTTP protocol
		if response := strings.TrimSpace(w.Body.String()); response != tt.response {
			t.Errorf("#%d HTTP status got %v, want %s", i, response, tt.response)
		}
	}
}

func TestGetNextFibonacciNumberFunc(t *testing.T) {
	var testcases = []struct {
		getFibonacciNumber getFibonacciNumberFunc

		response   string
		statusCode int
	}{
		// result response
		{
			getFibonacciNumber: func(ctx context.Context) (int64, error) {
				return 1, nil
			},
			response:   `{"next":1}`,
			statusCode: http.StatusOK,
		},
		// overflow error
		{
			getFibonacciNumber: func(ctx context.Context) (int64, error) {
				return 0, fibonacci.ErrCounterOverflow
			},
			response:   `{"error":"counter overflow"}`,
			statusCode: http.StatusInternalServerError,
		},
		// service error
		{
			getFibonacciNumber: func(ctx context.Context) (int64, error) {
				return 0, errors.New("test getFibonacciNumber error")
			},
			response:   "",
			statusCode: http.StatusInternalServerError,
		},
	}

	for i, tt := range testcases {
		req := httptest.NewRequest(http.MethodPost, "http://localhost/next", nil)
		w := httptest.NewRecorder()

		// To add the vars to the context we need to create a router through which we can pass the request.
		// TODO: tests should be not aware of routing mechanism?
		router := mux.NewRouter()
		router.HandleFunc("/next", GetNextFibonacciNumberFunc(tt.getFibonacciNumber))

		router.ServeHTTP(w, req)

		if statusCode := w.Result().StatusCode; statusCode != tt.statusCode {
			t.Errorf("#%d HTTP status got %v, want %v", i, statusCode, tt.statusCode)
		}

		// we do apply TrimSpace to clean up response coming from HTTP protocol
		if response := strings.TrimSpace(w.Body.String()); response != tt.response {
			t.Errorf("#%d HTTP status got %v, want %s", i, response, tt.response)
		}
	}
}

func TestGetPreviousFibonacciNumberFunc(t *testing.T) {
	var testcases = []struct {
		getFibonacciNumber getFibonacciNumberFunc

		response   string
		statusCode int
	}{
		// result response
		{
			getFibonacciNumber: func(ctx context.Context) (int64, error) {
				return 1, nil
			},
			response:   `{"previous":1}`,
			statusCode: http.StatusOK,
		},
		// underflow error
		{
			getFibonacciNumber: func(ctx context.Context) (int64, error) {
				return 0, fibonacci.ErrCounterUnderflow
			},
			response:   `{"error":"counter underflow"}`,
			statusCode: http.StatusInternalServerError,
		},
		// service error
		{
			getFibonacciNumber: func(ctx context.Context) (int64, error) {
				return 0, errors.New("test getFibonacciNumber error")
			},
			response:   "",
			statusCode: http.StatusInternalServerError,
		},
	}

	for i, tt := range testcases {
		req := httptest.NewRequest(http.MethodPost, "http://localhost/previous", nil)
		w := httptest.NewRecorder()

		// To add the vars to the context we need to create a router through which we can pass the request.
		// TODO: tests should be not aware of routing mechanism?
		router := mux.NewRouter()
		router.HandleFunc("/previous", GetPreviousFibonacciNumberFunc(tt.getFibonacciNumber))

		router.ServeHTTP(w, req)

		if statusCode := w.Result().StatusCode; statusCode != tt.statusCode {
			t.Errorf("#%d HTTP status got %v, want %v", i, statusCode, tt.statusCode)
		}

		// we do apply TrimSpace to clean up response coming from HTTP protocol
		if response := strings.TrimSpace(w.Body.String()); response != tt.response {
			t.Errorf("#%d HTTP status got %v, want %s", i, response, tt.response)
		}
	}
}
