package recursion1

import "strings"

func countHi(str string) int {
	if len(str) <= 1 {
		return 0
	}
	if strings.Compare(str[0:2], "hi") == 0 {
		return 1 + countHi(str[2:])
	}
	return countHi(str[1:])
}
