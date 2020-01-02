package ap1

func scoresClump(scores []int) bool {
	for i := 0; i < len(scores)-2; i++ {
		if scores[i+2]-scores[i] <= 2 {
			return true
		}
	}
	return false
}
