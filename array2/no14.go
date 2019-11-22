package array2

func no14(nums []int) bool {
	f1 := true
	f4 := true
	for _, n := range nums {
		if n == 1 {
			f1 = false
		}
		if n == 4 {
			f4 = false
		}
	}
	return f1 || f4
}
