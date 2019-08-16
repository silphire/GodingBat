package array1

func unlucky1(nums []int) bool {
	n := len(nums)
	return nums[0] == 1 && nums[1] == 3 || nums[n - 2] == 1 && nums[n - 1] == 3
}