package warmup1

import "strings"

// https://codingbat.com/prob/p191914

func notString(str string) string {
	if strings.HasPrefix(str, "not") {
		return str
	}
	return "not " + str
}
