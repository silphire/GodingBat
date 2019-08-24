package array1

func start1(a []int, b []int) int {
	n := 0
	if len(a) > 0 && a[0] == 1 {
		n++
	}
	if len(b) > 0 && b[0] == 1 {
		n++
	}
	return n
}