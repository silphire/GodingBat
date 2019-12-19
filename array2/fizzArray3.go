package array2

func fizzArray3(start int, end int) []int {
	n := end - start
	r := make([]int, n, n)
	for i := 0; i < n; i++ {
		r[i] = i + start
	}
	return r
}
