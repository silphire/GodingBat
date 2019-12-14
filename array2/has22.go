package array2

func has22(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 2 && nums[i+1] == 2 {
			return true
		}
	}
	return false
}
