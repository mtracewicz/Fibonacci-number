package fib

import "errors"

func RecursiveFib(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("must me positive integer")
	}

	if n == 0 || n == 1 {
		return n, nil
	} else {
		f, _ := RecursiveFib(n - 1)
		s, _ := RecursiveFib(n - 2)
		return f + s, nil
	}
}
