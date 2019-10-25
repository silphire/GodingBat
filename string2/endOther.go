package string2

import "strings"

func endOther(a string, b string) bool {
	la := strings.ToLower(a)
	lb := strings.ToLower(b)
	return strings.HasSuffix(la, lb) || strings.HasSuffix(lb, la)
}
