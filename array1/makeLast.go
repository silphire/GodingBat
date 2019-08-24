package array1

func makeLast(nums []int) []int {
	n := len(nums)
	r := make([]int, n * 2, n * 2)
	r[n * 2 - 1] = nums[n - 1]
	return r
}