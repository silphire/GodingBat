package warmup2

func noTriples(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	prev := nums[0]
	cnt := 0

	for i := 1; i < len(nums); i++ {
		if prev == nums[i] {
			cnt++
			if cnt >= 3 {
				return true
			}
		} else {
			cnt = 1
			prev = nums[i]
		}
	}

	return false
}
