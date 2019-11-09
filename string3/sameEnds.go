package string3

import "strings"

func sameEnds(str string) string {
	for i := len(str) / 2; i > 0; i-- {
		sl := string(str[0:i])
		sr := string(str[len(str)-i : len(str)])
		if strings.Compare(sl, sr) == 0 {
			return sl
		}
	}
	return ""
}
