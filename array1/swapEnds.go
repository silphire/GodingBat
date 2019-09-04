package array1

func swapEnds(nums []int) []int {
	n := len(nums)
	x := nums[n - 1]
	nums[n - 1] = nums[0]
	nums[0] = x
	return nums
}