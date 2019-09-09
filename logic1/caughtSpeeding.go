package logic1

func caughtSpeeding(speed int, isBirthday bool) int {
	if isBirthday {
		speed -= 5
	}
	if speed <= 60 {
		return 0
	}
	if speed <= 80 {
		return 1
	}
	return 2
}