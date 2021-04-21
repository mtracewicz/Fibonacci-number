package fib

func SlicesFib(n int) int {
	var tab = []int{0, 1}
	for i := 2; i <= n; i++ {
		tab = append(tab, tab[i-1]+tab[i-2])
	}
	return tab[n]
}
