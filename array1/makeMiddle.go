package array1

func makeMiddle(nums []int) []int {
	n := len(nums)
	return []int{nums[n / 2 - 1], nums[n / 2]}
}