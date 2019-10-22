package string2

import "strings"

func prefixAgain(str string, n int) bool {
	prefix := string(str[0:n])
	return strings.Count(str, prefix) > 1
}
