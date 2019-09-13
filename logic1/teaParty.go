package logic1

func teaParty(tea int, candy int) int {
	if tea >= 5 && candy >= 5 {
		if tea >= 10 || candy >= 10 {
			return 2
		}
		return 1
	}
	return 0
}