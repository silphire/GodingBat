package array2

func modThree(nums []int) bool {
	ne := 0
	no := 0
	for _, n := range nums {
		if n%2 == 0 {
			ne++
		} else {
			no++
		}
	}
	return ne == 3 || no == 3
}
