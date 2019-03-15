package warmup1

// https://codingbat.com/prob/p177372

func max1020(a int, b int) int {
	x := make([]int, 2)

	if a > b {
		x[0] = a
		x[1] = b
	} else {
		x[0] = b
		x[1] = a
	}

	for _, n := range x {
		if n >= 10 && n <= 20 {
			return n
		}
	}

	return 0
}
