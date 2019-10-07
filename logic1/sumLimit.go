package logic1

import "strconv"

func sumLimit(a int, b int) int {
	n := a + b
	if len(strconv.Itoa(n)) == len(strconv.Itoa(a)) {
		return n
	}
	return a
}
