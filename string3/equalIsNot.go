package string3

import "strings"

func equalIsNot(str string) bool {
	return strings.Count(str, "is") == strings.Count(str, "not")
}
