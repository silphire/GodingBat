package array2

func has77(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] != 7 {
			continue
		}
		if i+1 < len(nums) && nums[i+1] == 7 {
			return true
		}
		if i+2 < len(nums) && nums[i+2] == 7 {
			return true
		}
	}
	return false
}
