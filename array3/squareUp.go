package array3

func squareUp(n int) []int {
	sz := n * n
	r := make([]int, sz, sz)
	for i := 0; i < n; i++ {
		st := n - i - 1
		for j := 0; j < n; j++ {
			x := 0
			if j >= st {
				x = n - j
			}
			r[i*n+j] = x
		}
	}
	return r
}
