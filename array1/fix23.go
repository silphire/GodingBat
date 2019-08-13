package array1

func fix23(nums []int) []int {
	if nums[0] == 2 && nums[1] == 3 {
		return []int{nums[0], 0, nums[2]}
	}
	if nums[1] == 2 && nums[2] == 3 {
		return []int{nums[0], nums[1], 0}
	}
	return nums
}