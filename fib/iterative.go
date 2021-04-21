package fib

import "errors"

func IterativeFib(n int) (int, error) {
	var fib int
	if n < 0 {
		return 0, errors.New("must me positive integer")
	}

	x := 0
	y := 1
	z := 1
	for i := 0; i < (n / 3); i++ {
		x = y + z
		y = x + z
		z = y + x
	}
	if n%3 == 0 {
		fib = x
	} else if n%3 == 1 {
		fib = y
	} else {
		fib = z
	}
	return fib, nil
}
