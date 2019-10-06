package logic1

func withoutDoubles(die1 int, die2 int, noDoubles bool) int {
	if die1 == die2 && noDoubles {
		return 1 + die1 + die2
	}
	return die1 + die2
}
