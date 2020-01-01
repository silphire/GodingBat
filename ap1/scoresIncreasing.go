package ap1

func scoresIncreasing(scores []int) bool {
	for i := 1; i < len(scores); i++ {
		if scores[i-1] > scores[i] {
			return false
		}
	}
	return true
}
