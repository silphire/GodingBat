package array2

func zeroMax(nums []int) []int {
	n := len(nums)
	r := make([]int, n, n)
	maxv := nums[n-1]
	for i := n - 1; i >= 0; i-- {
		if nums[i] == 0 {
			r[i] = maxv
		} else {
			r[i] = nums[i]
		}
		if nums[i]%2 == 1 && maxv < nums[i] {
			maxv = nums[i]
		}
	}
	return r
}
