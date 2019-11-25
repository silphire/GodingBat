package array2

func sameEnds(nums []int, numslen int) bool {
	for i := 0; i < numslen; i++ {
		if nums[i] != nums[len(nums)-numslen+i] {
			return false
		}
	}
	return true
}
