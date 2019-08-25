package array1

func plusTwo(a []int, b []int) []int {
	return append(a, b...)
}