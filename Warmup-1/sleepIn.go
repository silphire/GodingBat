// https://codingbat.com/prob/p187868

package main

func sleepIn(weekday bool, vacation bool) bool {
	return !weekday || vacation
}

func action(doing bool) string {
	if doing {
		return "wakeup"
	} else {
		return "sleep"
	}
}