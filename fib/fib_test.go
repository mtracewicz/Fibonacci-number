package fib

import "testing"

func TestIterative(t *testing.T) {
	got, err := IterativeFib(20)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestIterativeLesserThen2(t *testing.T) {
	got, err := IterativeFib(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got != 1 {
		t.Errorf("Fib(1) = %d; want 1", got)
	}
}

func TestIterativeNegativeInteger(t *testing.T) {
	got, err := IterativeFib(-1)
	if err == nil || got != 0 {
		t.Errorf("Got value when error should be returned")
	}
}

func TestIterativeFibV2(t *testing.T) {
	got, err := IterativeFibV2(20)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestIterativeFibV2LesserThen2(t *testing.T) {
	got, err := IterativeFibV2(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got != 1 {
		t.Errorf("Fib(1) = %d; want 1", got)
	}
}

func TestIterativeFibV2NegativeInteger(t *testing.T) {
	got, err := IterativeFibV2(-1)
	if err == nil || got != 0 {
		t.Errorf("Got value when error should be returned")
	}
}

func TestRecursive(t *testing.T) {
	got, err := RecursiveFib(20)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestRecursiveLesserThen2(t *testing.T) {
	got, err := RecursiveFib(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got != 1 {
		t.Errorf("Fib(1) = %d; want 1", got)
	}
}

func TestRecursiveNegativeInteger(t *testing.T) {
	got, err := RecursiveFib(-1)
	if err == nil || got != 0 {
		t.Errorf("Got value when error should be returned")
	}
}

func TestSlices(t *testing.T) {
	got, err := SlicesFib(20)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got != 6765 {
		t.Errorf("Fib(20) = %d; want 6765", got)
	}
}

func TestSlicesLesserThen2(t *testing.T) {
	got, err := SlicesFib(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	if got != 1 {
		t.Errorf("Fib(1) = %d; want 1", got)
	}
}

func TestSlicesNegativeInteger(t *testing.T) {
	got, err := SlicesFib(-1)
	if err == nil || got != 0 {
		t.Errorf("Got value when error should be returned")
	}
}
