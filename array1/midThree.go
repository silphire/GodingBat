package array1

func midThree(nums []int) []int {
	n := len(nums) / 2
	return []int{nums[n - 1], nums[n], nums[n + 1]}
}