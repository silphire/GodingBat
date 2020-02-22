package recursion1

import "strings"

func strCopies(str string, sub string, n int) bool {
	if len(str) < len(sub) {
		return n == 0
	}
	if strings.Compare(str[0:len(sub)], sub) == 0 {
		return strCopies(str[len(sub):len(str)], sub, n-1)
	} else {
		return strCopies(str[1:len(str)], sub, n)
	}
}
