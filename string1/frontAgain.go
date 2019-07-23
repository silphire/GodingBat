package string1

import "strings"

func frontAgain(str string) bool {
	n := len(str)
	return strings.Compare(str[0:2], str[n-2:n]) == 0
}