package logic1

func teenSum(a int, b int) int {
	if a >= 13 && a <= 19 || b >= 13 && b <= 19 {
		return 19
	}
	return a + b
}
