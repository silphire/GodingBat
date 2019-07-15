package string1

import "strings"

func startWord(str string, word string) string {
	if strings.HasPrefix(str[1:], word[1:]) {
		return str[:len(word)]
	}
	return ""
}