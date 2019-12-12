package array2

import "sort"

func centeredAverage(nums []int) int {
	sort.Ints(nums)
	cnt := map[int]int{}
	for i := 1; i < len(nums)-1; i++ {
		_, ok := cnt[nums[i]]
		if ok {
			cnt[nums[i]]++
		} else {
			cnt[nums[i]] = 1
		}
	}

	s := 0
	n := 0
	for k, v := range cnt {
		if v > 1 {
			v--
		}
		s += k * v
		n += v
	}

	return s / n
}
