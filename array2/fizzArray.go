package array2

func fizzArray(n int) []int {
	r := make([]int, n, n)
	for i := 0; i < n; i++ {
		r[i] = i
	}
	return r
}
