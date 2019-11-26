package array2

func shiftLeft(nums []int) []int {
	r := make([]int, len(nums), len(nums))
	n := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		r[i] = nums[i+1]
	}
	r[len(r)-1] = n
	return r
}
