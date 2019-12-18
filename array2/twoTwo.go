package array2

func twoTwo(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 2 {
			if i+1 >= len(nums) || nums[i+1] != 2 {
				return false
			}
			i++
		}
	}
	return true
}
