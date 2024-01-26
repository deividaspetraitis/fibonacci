package fibonacci

import (
	"context"
	"sync"
	"testing"

	"github.com/deividaspetraitis/fibonacci/errors"
)

// direction defines direction for the Fibonacci to walk.
type direction func(ctx context.Context) (int64, error)

// testWalkFibonacci implements main logic for testing walking through the sequence.

// Parameter walk specifies which direction to walk through the sequence.
// Parameter expectedCount specifies how far to walk to given direction.
// Parameter expectedNumber specifies expected final sequence term+1 value.
func testWalkFibonacci(t *testing.T, sequence *Fibonacci, walk direction, expectedCount, expectedNumber int) {
	t.Helper()

	var wg sync.WaitGroup
	wg.Add(expectedCount)

	for i := 0; i < expectedCount; i++ {
		go func() {
			walk(context.TODO())
			wg.Done()
		}()
	}
	wg.Wait()

	if sequence.counter != expectedCount {
		t.Errorf("got %v, want %v", sequence.counter, expectedCount)
	}

	v, err := walk(context.TODO())
	if err != nil {
		t.Errorf("got %v, want %v", err, nil)
	}

	if v != int64(expectedNumber) {
		t.Errorf("got %v, want %v", v, expectedNumber)
	}
}

func TestCurrentFibonacciNumber(t *testing.T) {
	var sequence Fibonacci
	testWalkFibonacci(t, &sequence, (&sequence).NextFibonacciNumber, 60, 2504730781961)

	var testcases = []struct {
		sequence *Fibonacci

		expected int64
	}{
		{
			sequence: &Fibonacci{counter: 73},
			expected: 806515533049393,
		},
		{
			sequence: &Fibonacci{counter: 44},
			expected: 701408733,
		},
		{

			sequence: &Fibonacci{counter: 88},
			expected: 1100087778366101931,
		},
		{
			sequence: &Fibonacci{counter: 92},
			expected: 7540113804746346429,
		},
		{
			sequence: &Fibonacci{counter: 1},
			expected: 1,
		},
	}

	for _, tt := range testcases {
		got, err := tt.sequence.CurrentFibonacciNumber(context.TODO())
		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}

		if got != tt.expected {
			t.Errorf("#%dth got %v, want %v", tt.sequence.counter, got, tt.expected)
		}
	}
}

func TestPreviousFibonacciNumber(t *testing.T) {
	sequence := Fibonacci{
		counter: 30,
	}
	testWalkFibonacci(t, &sequence, (&sequence).PreviousFibonacciNumber, 15, 377)
}

func TestNextFibonacciNumber(t *testing.T) {
	var sequence Fibonacci
	testWalkFibonacci(t, &sequence, (&sequence).NextFibonacciNumber, 60, 2504730781961)

	// test overflow error
	sequence.counter = MaxThTerm
	if _, err := (&sequence).NextFibonacciNumber(context.TODO()); !errors.Is(err, ErrCounterOverflow) {
		t.Errorf("got %v, want %v", err, ErrCounterOverflow)
	}
}

func TestCalcFiboncciTerm(t *testing.T) {
	var testcases = []struct {
		n int

		expected int64
	}{
		{
			n:        73,
			expected: 806515533049393,
		},
		{
			n:        44,
			expected: 701408733,
		},
		{
			n:        88,
			expected: 1100087778366101931,
		},
		{
			n:        92,
			expected: 7540113804746346429,
		},
		{
			n:        1,
			expected: 1,
		},
	}

	for _, tt := range testcases {
		got := calcFiboncciTerm(tt.n)
		if got != tt.expected {
			t.Errorf("#%dth got %v, want %v", tt.n, got, tt.expected)
		}
	}
}
