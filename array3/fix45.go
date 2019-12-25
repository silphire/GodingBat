package array3

func fix45(nums []int) []int {
	n := len(nums)
	r := make([]int, n, n)
	for i := 0; i < n; i++ {
		if nums[i] == 4 {
			r[i] = 4
			i++
			r[i] = 5
		} else {
			r[i] = 0
		}
	}

	j := 0
	for i := 0; i < n; i++ {
		if nums[i] == 4 || nums[i] == 5 {
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
