package array1

func make2(a []int, b []int) []int {
	j := 0
	r := make([]int, 2, 2)
	for i := 0; i < len(a); i++ {
		r[j] = a[i]
		j++
		if j == 2 {
			return r
		}
	}
	for i := 0; i < len(b); i++ {
		r[j] = b[i]
		j++
		if j == 2 {
			return r
		}
	}
	return r
}