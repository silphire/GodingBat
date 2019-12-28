package array3

func seriesUp(n int) []int {
	sz := (n + 1) * n / 2
	r := make([]int, sz, sz)

	k := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			r[k] = j
			k++
		}
	}

	return r
}
