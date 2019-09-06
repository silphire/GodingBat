package logic1

func cigarParty(cigars int, isWeekend bool) bool {
	if isWeekend {
		return true
	}
	return cigars >= 40 && cigars <= 60
}