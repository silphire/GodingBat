package ap1

func scoresAverage(scores []int) int {
	n := len(scores) / 2
	a := 0
	b := 0
	for i := 0; i < n; i++ {
		a += scores[i]
		b += scores[i+n]
	}
	if a > b {
		return a / n
	}
	return b / n
}
