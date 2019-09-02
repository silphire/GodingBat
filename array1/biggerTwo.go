package array1

func biggerTwo(a []int, b []int) []int {
	if a[0] + a[1] > b[0] + b[1] {
		return a
	} else {
		return b
	}
}