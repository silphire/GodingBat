package array2

func countEvents(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			n++
		}
	}
	return n
}
