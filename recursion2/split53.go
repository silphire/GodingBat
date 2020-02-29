package recursion2

func sub(nums []int, i int, left int, right int) bool {
	if i == len(nums) {
		return left == right
	}
	if nums[i]%5 == 0 {
		return sub(nums, i+1, left+nums[i], right)
	}
	if nums[i]%3 == 0 {
		return sub(nums, i+1, left, right+nums[i])
	}
	if sub(nums, i+1, left+nums[i], right) {
		return true
	}
	return sub(nums, i+1, left, nums[i]+right)
}

func split53(nums []int) bool {
	return sub(nums, 0, 0, 0)
}
