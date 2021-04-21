package fib

import "errors"

func SlicesFib(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("must me positive integer")
	}

	var tab = []int{0, 1}
	for i := 2; i <= n; i++ {
		tab = append(tab, tab[i-1]+tab[i-2])
	}
	return tab[n], nil
}
