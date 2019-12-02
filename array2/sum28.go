package array2

func sum28(nums []int) bool {
	n := 0
	for _, x := range nums {
		if x == 2 {
			n += x
		}
	}
	return n == 8
}
