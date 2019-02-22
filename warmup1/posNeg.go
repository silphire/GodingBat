package warmup1

// https://codingbat.com/prob/p159227

func posNeg(a int, b int, negative bool) bool {
	if negative {
		return a < 0 && b < 0
	}
	return a < 0 && b >= 0 || a >= 0 && b < 0
}
