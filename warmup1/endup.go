package warmup1

import "strings"

func endUp(str string) string {
	n := len(str)
	if n <= 3 {
		return strings.ToUpper(str)
	}
	return str[0:n-3] + strings.ToUpper(str[n-3:n])
}
