package array2

import "sort"

func zeroFront(nums []int) []int {
	sort.Ints(nums[:])
	return nums
}
