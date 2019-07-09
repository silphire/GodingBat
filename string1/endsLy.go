package string1

import "strings"

func endsLy(str string) bool {
	n := len(str)
	return n >= 2 && strings.Compare(str[n-2:n], "ly") == 0
}
