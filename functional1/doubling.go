package functional1

import "github.com/thoas/go-funk"

func doubling(nums []int) []int {
	return funk.Map(nums, func(n int) int {
		return n * 2
	}).([]int)
}
