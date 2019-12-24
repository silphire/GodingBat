package array3

func fix34(nums []int) []int {
	n := len(nums)
	r := make([]int, n, n)
	for i := 0; i < n; i++ {
		if nums[i] == 3 {
			r[i] = 3
			i++
			r[i] = 4
		} else {
			r[i] = 0
		}
	}

	j := 0
	for i := 0; i < n; i++ {
		if nums[i] == 3 || nums[i] == 4 {
			continue
		}
		for r[j] != 0 {
			j++
		}
		r[j] = nums[i]
		j++
	}
	return r
}
