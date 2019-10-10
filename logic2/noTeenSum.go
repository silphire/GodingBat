package logic2

func fixTeen(n int) int {
	if n == 15 || n == 16 {
		return n
	}
	if n >= 13 && n <= 19 {
		return 0
	}
	return n
}

func noTeenSum(a int, b int, c int) int {
	return fixTeen(a) + fixTeen(b) + fixTeen(c)
}
