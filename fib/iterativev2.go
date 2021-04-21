package fib

import "errors"

func IterativeFibV2(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("must me positive integer")
	}

	x := 0
	y := 1
	for i := 0; i < n; i++ {
		y += x
		x = y - x
	}
	return x, nil
}
