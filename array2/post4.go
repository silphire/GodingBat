package array2

func post4(nums []int) []int {
	i := 0
	for i = len(nums) - 1; i >= 0; i-- {
		if nums[i] == 4 {
			break
		}
	}
	return nums[i+1 : len(nums)]
}
