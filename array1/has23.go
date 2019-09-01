package array1

func has23(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 2 || nums[i] == 3 {
			return true
		}
	}
	return false
}