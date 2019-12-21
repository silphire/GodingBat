package array2

func pre4(nums []int) []int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 4 {
			n = i
			break
		}
	}
	r := make([]int, n, n)
	for i := 0; i < n; i++ {
		r[i] = nums[i]
	}
	return r
}
