package main

func iterativeFib(n int) (fib int) {
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
	return
}
