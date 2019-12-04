package array2

func only14(nums []int) bool {
	for _, n := range nums {
		if n != 1 && n != 4 {
			return false
		}
	}
	return true
}
