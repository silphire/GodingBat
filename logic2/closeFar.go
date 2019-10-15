package logic2

func abs(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func closeFar(a int, b int, c int) bool {
	return abs(a, b) == 1 && abs(a, c) > 1 && abs(b, c) > 1 || abs(a, c) == 1 && abs(a, b) > 1 && abs(b, c) > 1
}
