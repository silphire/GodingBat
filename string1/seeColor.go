package string1

import "strings"

func seeColor(str string) string {
	if strings.HasPrefix(str, "red") {
		return "red"
	} else if strings.HasPrefix(str, "blue") {
		return "blue"
	} else {
		return ""
	}
}