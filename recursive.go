package main

func recursiveFib(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return recursiveFib(n-1) + recursiveFib(n-2)
	}
}
