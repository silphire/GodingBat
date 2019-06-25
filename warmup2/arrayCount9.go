package warmup2

func arrayCount9(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 9 {
			n++
		}
	}
	return n
}
