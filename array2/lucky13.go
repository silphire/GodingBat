package array2

func lucky13(nums []int) bool {
	for _, n := range nums {
		if n == 1 || n == 3 {
			return false
		}
	}
	return true
}
