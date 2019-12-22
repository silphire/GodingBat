package array2

import "sort"

func evenOdd(nums []int) []int {
	sort.Slice(
		nums,
		func(i int, j int) bool {
			return (nums[i] & 1) < (nums[j] & 1)
		},
	)
	return nums
}
