package logic2

func makeBricks(small int, big int, goal int) bool {
	if big*5+small < goal {
		return false
	}
	if goal-5*big < 0 {
		if small < goal%5 {
			return false
		}
	} else {
		return small >= goal%5
	}
	return true
}
