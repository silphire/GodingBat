package array3

func countClumps(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	r := 1
	l := 1
	p := nums[0]
	for i := 1; i < len(nums); i++ {
		if p == nums[i] {
			l++
		} else {
			if l > 1 {
				r++
			}
			l = 1
			p = nums[i]
		}
	}
	return r
}
