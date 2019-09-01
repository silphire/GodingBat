package array1

func double23(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	return (nums[0] == 2 && nums[1] == 2) || (nums[0] == 3 && nums[1] == 3)
}