package logic2

func blackjack(a int, b int) int {
	if a > 21 && b > 21 {
		return 0
	}
	if a > 21 {
		return b
	}
	if b > 21 {
		return a
	}
	if a < b {
		return b
	}
	return a
}
