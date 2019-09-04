package array1

func frontPiece(nums []int) []int {
	x := make([]int, 2, 2)
	n := len(nums)
	if n > 2 {
		n = 2
	}
	for i := 0; i < n; i++ {
		x[i] = nums[i]
	}
	return x
}