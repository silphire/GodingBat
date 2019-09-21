package logic1

func maxMod5(a int, b int) int {
	if a%5 == b%5 {
		if a > b {
			return b
		}
		return a
	}
	if a > b {
		return a
	}
	return b
}
