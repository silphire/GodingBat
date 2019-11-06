package string2

import "strings"

func wordEnds(str string, word string) string {
	r := ""
	ss := strings.Split(str, word)
	for i := 0; i < len(ss); i++ {
		if len(ss[i]) == 0 {
			continue
		}
		if i > 0 {
			r += string(ss[i][0])
		}
		if i < len(ss)-1 {
			r += string(ss[i][len(ss[i])-1])
		}
	}
	return r
}
