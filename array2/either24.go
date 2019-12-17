package array2

func either24(nums []int) bool {
	f2 := false
	f4 := false
	for i := 0; i < len(nums)-1; i++ {
		f2 = f2 || nums[i] == 2 && nums[i+1] == 2
		f4 = f4 || nums[i] == 4 && nums[i+1] == 4
	}
	return f2 && !f4 || !f2 && f4
}
