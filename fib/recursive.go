package fib

func RecursiveFib(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return RecursiveFib(n-1) + RecursiveFib(n-2)
	}
}
