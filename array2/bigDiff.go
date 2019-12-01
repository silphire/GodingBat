package array2

func bigDiff(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ma := nums[0]
	mi := nums[0]
	for i := 1; i < len(nums); i++ {
		if ma < nums[i] {
			ma = nums[i]
		}
		if mi > nums[i] {
			mi = nums[i]
		}
	}
	return ma - mi
}
