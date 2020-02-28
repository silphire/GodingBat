package recursion2

func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sum(nums[1:len(nums)])
}

func make10(nums []int, remain int) bool {
	if remain < 0 || len(nums) == 0 {
		return false
	}
	if remain == 0 {
		return sum(nums)%2 == 1
	}
	return make10(nums[1:len(nums)], remain-nums[0])

}

func splitOdd10(nums []int) bool {
	return make10(nums, 10)
}
