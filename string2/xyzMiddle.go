package string2

import "strings"

func xyzMiddle(str string) bool {
	r := strings.Split(str, "xyz")
	if len(r) != 2 {
		return false
	}
	return len(r[0]) == len(r[1]) || len(r[0])-len(r[1]) == 1 || len(r[1])-len(r[0]) == 1
}
