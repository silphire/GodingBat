package array2

func tripleUp(nums []int) bool {
	for i := 0; i < len(nums)-2; i++ {
		if nums[i]+2 == nums[i+1]+1 && nums[i+1]+1 == nums[i+2] {
			return true
		}
	}
	return false
}
