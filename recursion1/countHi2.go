package recursion1

import "strings"

func countHi2(str string) int {
	n := len(str)

	if n <= 1 {
		return 0
	}
	if n >= 3 && strings.Compare(str[0:3], "xhi") == 0 {
		return countHi2(str[2:n])
	}

	x := 0
	if strings.Compare(str[0:2], "hi") == 0 {
		x = 1
	}
	return x + countHi2(str[1:n])
}
