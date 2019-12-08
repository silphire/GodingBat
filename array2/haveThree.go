package array2

func haveThree(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 3 {
			cnt++
			if i > 0 && nums[i-1] == 3 {
				return false
			}
		}
	}
	return cnt == 3
}
