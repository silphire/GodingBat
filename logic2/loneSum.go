package logic2

func loneSum(a int, b int, c int) int {
	if a == b && b == c && c == a {
		return 0
	}
	if a == b {
		return c
	}
	if b == c {
		return a
	}
	if c == a {
		return b
	}
	return a + b + c
}
