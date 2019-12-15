package array2

func more14(nums []int) bool {
	n1 := 0
	n4 := 0
	for _, x := range nums {
		if x == 1 {
			n1++
		}
		if x == 4 {
			n4++
		}
	}
	return n1 > n4
}
