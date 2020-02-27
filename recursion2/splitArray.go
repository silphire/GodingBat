package recursion2

func _search(nums []int, target int) bool {
	if len(nums) == 0 || target < 0 {
		return false
	}
	if target == 0 {
		return true
	}
	return _search(nums[1:len(nums)], target-nums[0])
}

func _sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + _sum(nums[1:len(nums)])
}

func splitArray(nums []int) bool {
	x := _sum(nums)
	if x%2 == 1 {
		return false
	}
	x /= 2
	return _search(nums, x)
}
