package array1

func sameFirstLast(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	return nums[0] == nums[n - 1]
}