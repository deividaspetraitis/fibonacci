package api

import (
	"encoding/json"
	"net/http"
)

// Error represents an error response.
type Error struct {
	Message string `json:"error"`
}

// MarshalHTTP implements http.Marshaler.
func (r *Error) MarshalHTTP(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(r)
}

// CurrentFibonacciNumberResponse represents a response for getting current number in the Fibonacci sequence.
type CurrentFibonacciNumberResponse struct {
	Current int64 `json:"current"`
}

// MarshalHTTP implements http.Marshaler.
func (r *CurrentFibonacciNumberResponse) MarshalHTTP(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(r)
}

// NextFibonacciNumberResponse represents a response for getting next number in the Fibonacci sequence.
type NextFibonacciNumberResponse struct {
	Next int64 `json:"next"`
}

// MarshalHTTP implements http.Marshaler.
func (r *NextFibonacciNumberResponse) MarshalHTTP(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(r)
}

// PreviousFibonacciNumberResponse represents a response for getting previous number in the Fibonacci sequence.
type PreviousFibonacciNumberResponse struct {
	Previous int64 `json:"previous"`
}

// MarshalHTTP implements http.Marshaler.
func (r *PreviousFibonacciNumberResponse) MarshalHTTP(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(r)
}
