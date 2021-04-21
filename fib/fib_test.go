package fib

import "testing"

func TestIterative(t *testing.T) {
	got := IterativeFib(20)
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestIterativeFibV2(t *testing.T) {
	got := IterativeFibV2(20)
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestRecursive(t *testing.T) {
	got := RecursiveFib(20)
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestSlices(t *testing.T) {
	got := SlicesFib(20)
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}
