package recursion1

import "strings"

func count11(str string) int {
	if len(str) < 2 {
		return 0
	}
	if strings.Compare(str[0:2], "11") == 0 {
		return 1 + count11(str[2:len(str)])
	}
	return count11(str[1:len(str)])
}
