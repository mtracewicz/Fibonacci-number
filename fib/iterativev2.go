package fib

func IterativeFibV2(n int) int {
	x := 0
	y := 1
	for i := 0; i < n; i++ {
		y += x
		x = y - x
	}
	return x
}
