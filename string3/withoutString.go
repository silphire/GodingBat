package string3

import "strings"

func withoutString(base string, remove string) string {
	return strings.Replace(base, remove, "", -1)
}
