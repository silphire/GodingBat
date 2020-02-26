package recursion2

func groupSumClump(start int, nums []int, target int) bool {
	if start > len(nums) || target < 0 {
		return false
	}
	if start == len(nums) {
		return target == 0
	}

	n := 0
	for start+n < len(nums) && nums[start+n] == nums[start] {
		n++
	}
	if groupSumClump(start+n, nums, target-nums[start]*n) {
		return true
	}
	return groupSumClump(start+n, nums, target)
}
