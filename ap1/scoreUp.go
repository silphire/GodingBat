package ap1

func scoreUp(key []string, answers []string) int {
	score := 0
	for i := 0; i < len(key); i++ {
		if answers[i] == "?" {
			continue
		}
		if answers[i] == key[i] {
			score += 4
		} else {
			score--
		}
	}
	return score
}
