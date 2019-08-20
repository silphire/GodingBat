package array1

func sum3(nums []int) int {
	n := len(nums)
	if n > 3 {
		n = 3
	}

	x := 0
	for i := 0; i < n; i++ {
		x += nums[i]
	}
	return x
}