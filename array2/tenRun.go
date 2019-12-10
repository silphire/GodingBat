package array2

func tenRun(nums []int) []int {
	n := len(nums)
	r := make([]int, n, n)
	x := 0
	for i := 0; i < n; i++ {
		if nums[i]%10 == 0 {
			x = nums[i]
		}
		if x == 0 {
			r[i] = nums[i]
		} else {
			r[i] = x
		}
	}
	return r
}
