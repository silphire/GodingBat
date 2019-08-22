package array1

func makeEnds(nums []int) []int {
	n := len(nums)
	return []int{nums[0], nums[n - 1]}
}