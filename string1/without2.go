package string1

import "strings"

func without2(str string) string {
	n := len(str)
	if n >= 2 && strings.Compare(str[0:2], str[n-2:n]) == 0 {
		return str[2:n]
	}
	return str
}