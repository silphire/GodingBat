package recursion1

import "strings"

func strCount(str string, sub string) int {
	if len(str) < len(sub) {
		return 0
	}
	if strings.Compare(str[0:len(sub)], sub) == 0 {
		return 1 + strCount(str[len(sub):len(str)], sub)
	}
	return strCount(str[1:len(str)], sub)
}
