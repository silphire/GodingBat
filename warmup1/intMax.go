package warmup1

// https://codingbat.com/prob/p101887

func intMax(a int, b int, c int) int {
	if a > b {
		if c > a {
			return c
		}
		return a
	}
	if c > b {
		return c
	}
	return b
}
