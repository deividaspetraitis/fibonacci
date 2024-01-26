package fibonacci

import (
	"context"
	"math/big"
	"sync"

	"github.com/deividaspetraitis/fibonacci/errors"
)

// MaxThTerm defines a maximum th term computed for Fibonacci sequence.
const MaxThTerm = 92

// Fibonacci Errors
var (
	ErrCounterOverflow  = errors.New("fibonacci: next term overflows highest allowed term in the sequence")
	ErrCounterUnderflow = errors.New("fibonacci: next term underflows lowest allowed term in the sequence")
)

// Fibonacci implements walking through the sequence.
// It is safe to use Fibonacci concurrently.
type Fibonacci struct {
	// counter is safe to use concurrently.
	mu      sync.Mutex
	counter int
}

// CurrentFibonacciNumber returns the current number in the Fibonacci sequence.
func (f *Fibonacci) CurrentFibonacciNumber(ctx context.Context) (int64, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return calcFiboncciTerm(f.counter), nil
}

// GetNextFibonacciNumberFunc responds with the next number in the Fibonacci sequence.
func (f *Fibonacci) NextFibonacciNumber(ctx context.Context) (int64, error) {
	f.mu.Lock()
	// MaxThTerm term in the sequence is the largest to fix into uint.
	if (f.counter + 1) > MaxThTerm {
		defer f.mu.Unlock()
		return 0, ErrCounterOverflow
	} else {
		f.counter++
	}
	f.mu.Unlock()
	return calcFiboncciTerm(f.counter), nil
}

// GetPreviousFibonacciNumberFunc responds with the next number in the Fibonacci sequence.
func (f *Fibonacci) PreviousFibonacciNumber(ctx context.Context) (int64, error) {
	f.mu.Lock()
	if (f.counter - 1) < 0 {
		defer f.mu.Unlock()
		return 0, ErrCounterUnderflow
	}
	f.counter--
	f.mu.Unlock()
	return calcFiboncciTerm(f.counter), nil
}

// calcFiboncciTerm calculates and returns n th term of the Fibonacci sequence.
// This implementation is not efficient of O(n).
func calcFiboncciTerm(n int) int64 {
	f := big.NewInt(0)
	a, b := big.NewInt(0), big.NewInt(1)
	for i := 0; i <= n; i++ {
		f.Set(a)
		a.Set(b)
		b.Add(f, b)
	}
	return f.Int64()
}
