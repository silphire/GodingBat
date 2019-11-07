package string3

import "strings"

func countYZ(str string) int {
	ss := strings.Split(str, " ")
	r := 0
	for i := 0; i < len(ss); i++ {
		c := ss[i][len(ss[i])-1]
		if c == 'y' || c == 'z' {
			r++
		}
	}
	return r
}
