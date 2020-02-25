package recursion2

func groupNoAdj(start int, nums []int, target int) bool {
	if start == target {
		return true
	}
	if len(nums) == 0 {
		return false
	}
	if groupNoAdj(start, nums[1:len(nums)], target) {
		return true
	}

	if len(nums) <= 1 {
		return false
	}
	return groupNoAdj(start+nums[0], nums[2:len(nums)], target)
}
