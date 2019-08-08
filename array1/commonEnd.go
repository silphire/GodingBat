package array1

func commonEnd(a []int, b []int) bool {
	na := len(a)
	nb := len(b)
	if na == 0 || nb == 0 {
		return false
	}
	if a[0] != b[0] {
		return false
	}
	return a[na - 1] == b[nb - 1]
}