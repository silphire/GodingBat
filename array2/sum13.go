package array2

func sum13(nums []int) int {
	n := 0
	for _, x := range nums {
		if x != 13 {
			n += x
		}
	}
	return n
}
