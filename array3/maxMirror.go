package array3

func maxMirror(nums []int) int {
	ml := 0
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j >= i+1; j-- {
			if nums[i] == nums[j] {
				k := 0
				for ; i+k < len(nums) && j-k >= 0; k++ {
					if nums[i+k] != nums[j-k] {
						break
					}
				}
				if k > ml {
					ml = k
				}
			}
		}
	}
	return ml
}
