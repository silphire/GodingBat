package logic1

func alarmClock(day int, vacation bool) string {
	if vacation || day == 0 || day == 6 {
		return "10:00"
	}
	return "7:00"
}
