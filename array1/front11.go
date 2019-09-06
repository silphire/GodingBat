package array1

func front11(a []int, b []int) []int {
	if len(a) == 0 && len(b) == 0 {
		return []int{}
	}
	if len(a) == 0 && len(b) > 0 {
		return []int{b[0]}
	}
	if len(a) > 0 && len(b) == 0 {
		return []int{a[0]}
	}
	return []int{a[0], b[0]}
}