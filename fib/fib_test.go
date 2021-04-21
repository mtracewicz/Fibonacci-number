package fib

import "testing"

func TestIterative(t *testing.T) {
	got := IterativeFib(20)
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestIterativeLesserThen2(t *testing.T) {
	got := IterativeFib(1)
	if got != 1 {
		t.Errorf("Fib(1) = %d; want 1", got)
	}
}

func TestIterativeFibV2(t *testing.T) {
	got := IterativeFibV2(20)
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestIterativeFibV2LesserThen2(t *testing.T) {
	got := IterativeFibV2(1)
	if got != 1 {
		t.Errorf("Fib(1) = %d; want 1", got)
	}
}

func TestRecursive(t *testing.T) {
	got := RecursiveFib(20)
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestRecursiveLesserThen2(t *testing.T) {
	got := RecursiveFib(1)
	if got != 1 {
		t.Errorf("Fib(1) = %d; want 1", got)
	}
}

func TestSlices(t *testing.T) {
	got := SlicesFib(20)
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestSlicesLesserThen2(t *testing.T) {
	got := SlicesFib(1)
	if got != 1 {
		t.Errorf("Fib(1) = %d; want 1", got)
	}
}
