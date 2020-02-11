package recursion1

func array220(nums []int, index int) bool {
	if index+1 >= len(nums) {
		return false
	}
	if nums[index+1] == nums[index]*10 {
		return true
	}
	return array220(nums, index+1)
}
