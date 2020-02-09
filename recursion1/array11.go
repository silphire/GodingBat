package recursion1

func array11(nums []int, index int) int {
	if index == len(nums) {
		return 0
	}

	x := 0
	if nums[index] == 11 {
		x = 1
	}

	return x + array11(nums, index+1)
}
