package array1

func maxEnd3(nums []int) []int {
	n := nums[0]
	if nums[2] > nums[0] {
		n = nums[2]
	}
	return []int{n, n, n}
}