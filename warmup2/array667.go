package warmup2

func array667(nums []int) int {
	r := 0
	n := len(nums) - 1
	for i := 0; i < n; i++ {
		if nums[i] == 6 && (nums[i+1] == 6 || nums[i+1] == 7) {
			r++
		}
	}
	return r
}
