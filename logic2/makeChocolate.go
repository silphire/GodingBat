package logic2

func makeChocolate(small int, big int, goal int) int {
	if big*5+small < goal {
		return -1
	}
	if big*5 < goal {
		return goal - big*5
	}
	return goal % 5
}
