package array2

func withoutTen(nums []int) []int {
	r := make([]int, len(nums), len(nums))
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 10 {
			r[j] = nums[i]
			j++
		}
	}

	for j < len(nums) {
		r[j] = 0
		j++
	}
	return r
}
