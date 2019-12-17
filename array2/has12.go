package array2

func has12(nums []int) bool {
	f := false
	for _, n := range nums {
		if n == 1 {
			f = true
		}
		if f && n == 2 {
			return true
		}
	}
	return false
}
