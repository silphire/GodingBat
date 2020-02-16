package recursion1

import "strings"

func countAbc(str string) int {
	if len(str) < 3 {
		return 0
	}
	s := str[0:3]
	x := 0
	if strings.Compare(s, "abc") == 0 || strings.Compare(s, "aba") == 0 {
		x = 1
	}
	return x + countAbc(str[1:len(str)])
}
