package logic1

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func lessBy10(a int, b int, c int) bool {
	return abs(a-b) >= 10 || abs(b-c) >= 10 || abs(c-a) >= 10
}
