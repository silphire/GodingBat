package array2

func isEverywhere(nums []int, val int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != val && nums[i+1] != val {
			return false
		}
	}
	return true
}
