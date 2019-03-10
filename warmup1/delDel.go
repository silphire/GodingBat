package warmup1

import "strings"

// https://codingbat.com/prob/p100905

func delDel(str string) string {
	if strings.Index(str, "del") == 1 {
		return str[0:1] + str[4:]
	}
	return str
}
