// https://codingbat.com/prob/p187868

package main

import "fmt"

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

func main() {
	fmt.Printf("its weekday and its vacation then %s\n", action(sleepIn(true, true)))
	fmt.Printf("not weekday and its vacation then %s\n", action(sleepIn(false, true)))
	fmt.Printf("its weekday and not vacation then %s\n", action(sleepIn(true, false)))
	fmt.Printf("not weekday and not vacation then %s\n", action(sleepIn(false, false)))
}
