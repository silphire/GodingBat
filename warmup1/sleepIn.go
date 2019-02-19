// https://codingbat.com/prob/p187868

package warmup1

func sleepIn(weekday bool, vacation bool) bool {
	return !weekday || vacation
}

func action(doing bool) string {
	if doing {
		return "wakeup"
	}
	return "sleep"
}
