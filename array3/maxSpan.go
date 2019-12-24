package array3

func maxSpan(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := 1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				mm := j - i + 1
				if m < mm {
					m = mm
				}
			}
		}
	}
	return m
}