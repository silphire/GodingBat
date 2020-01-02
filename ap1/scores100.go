package ap1

func scores100(scores []int) bool {
	for i := 1; i < len(scores); i++ {
		if scores[i] == 100 && scores[i-1] == 100 {
			return true
		}
	}
	return false
}
