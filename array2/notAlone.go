package array2

func notAlone(nums []int, val int) []int {
	n := len(nums)
	r := make([]int, n, n)
	for i := 0; i < n; i++ {
		if i > 0 && i < len(nums)-1 {
			if nums[i] == val {
				x := nums[i-1]
				if x < nums[i+1] {
					x = nums[i+1]
				}
				r[i] = x
				continue
			}
		}
		r[i] = nums[i]
	}
	return r
}
