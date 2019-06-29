package warmup2

func has271(nums []int) bool {
	ans := []int{2, 7, 1}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == ans[j] {
			j++
			if j == len(ans) {
				return true
			}
		} else {
			j = 0
		}
	}
	return false
}
