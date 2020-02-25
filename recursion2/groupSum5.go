package recursion2

func groupSum5(start int, nums []int, target int) bool {
	if target < 0 {
		return false
	}
	if start == len(nums) {
		return target == 0
	}
	if nums[start]%5 == 0 {
		if start+1 < len(nums) && nums[start+1] == 1 {
			return groupSum5(start+2, nums, target-nums[start])
		}
		return groupSum5(start+1, nums, target-nums[start])
	}
	if groupSum5(start+1, nums, target) {
		return true
	}
	return groupSum5(start+1, nums, target-nums[start])
}
