package string3

import (
	"strconv"
	"unicode"
)

func sumDigits(str string) int {
	r := 0
	for _, c := range str {
		if unicode.IsDigit(c) {
			d, _ := strconv.Atoi(string(c))
			r += d
		}
	}
	return r
}
