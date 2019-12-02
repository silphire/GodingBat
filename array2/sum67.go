package array2

func sum67(nums []int) int {
	s := 0
	f := false
	for i := 0; i < len(nums); i++ {
		if f {
			if nums[i] == 7 {
				f = false
			}
		} else {
			if nums[i] == 6 {
				f = true
			} else {
				s += nums[i]
			}
		}
	}
	return s
}
