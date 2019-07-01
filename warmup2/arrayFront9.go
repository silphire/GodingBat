package warmup2

func arrayFront9(nums []int) bool {
	l := len(nums)
	if l > 4 {
		l = 4
	}
	for i := 0; i < l; i++ {
		if nums[i] == 9 {
			return true
		}
	}
	return false
}
