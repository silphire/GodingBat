package ap1

import "strings"

func wordsWithout(words []string, target string) []string {
	n := 0
	for i := 0; i < len(words); i++ {
		if strings.Compare(words[i], target) != 0 {
			n++
		}
	}

	r := make([]string, n, n)
	p := 0
	for i := 0; i < len(words); i++ {
		if strings.Compare(words[i], target) != 0 {
			r[p] = words[i]
			p++
		}
	}

	return r
}
