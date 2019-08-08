package array1

func firstLast6(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if nums[0] == 6 {
		return true
	}
	if nums[len(nums) - 1] == 6 {
		return true
	}
	return false
}