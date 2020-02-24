package recursion2

func groupSum6(start int, nums []int, target int) bool {
	if target < 0 {
		return false
	}
	if start == len(nums) {
		return target == 0
	}
	if nums[start] == 6 {
		return groupSum(start+1, nums, target-nums[start])
	}
	if groupSum6(start+1, nums, target) {
		return true
	}
	if groupSum6(start+1, nums, target-nums[start]) {
		return true
	}
	return false
}
