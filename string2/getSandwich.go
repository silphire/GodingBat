package string2

import "strings"

func getSandwich(str string) string {
	r := strings.Split(str, "bread")
	if len(r)%2 == 1 {
		return r[len(r)/2]
	}
	return ""
}
