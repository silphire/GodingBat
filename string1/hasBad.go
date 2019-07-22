package string1

import "strings"

func hasBad(str string) bool {
	return strings.Contains(str, "bad")
}