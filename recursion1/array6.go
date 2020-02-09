package recursion1

func array6(nums []int, index int) bool {
	if index == len(nums) {
		return false
	}
	if nums[index] == 6 {
		return true
	}
	return array6(nums, index+1)
}
