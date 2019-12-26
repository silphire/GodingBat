package array3

func canBalance(nums []int) bool {
	x := 0
	y := 0
	for _, a := range nums {
		x += a
	}
	for _, a := range nums {
		if x == y {
			return true
		}
		x -= a
		y += a
	}
	return false
}
