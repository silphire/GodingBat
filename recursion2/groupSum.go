package recursion2

func groupSum(start int, nums []int, target int) bool {
	if start == target {
		return true
	}
	if len(nums) == 0 {
		return false
	}
	if groupSum(start, nums[1:len(nums)], target) {
		return true
	}
	return groupSum(start+nums[0], nums[1:len(nums)], target)
}
