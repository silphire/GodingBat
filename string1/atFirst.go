package string1

func atFirst(str string) string {
	n := len(str)
	if n == 0 {
		return "@@"
	} else if n == 1 {
		return str[0:1] + "@"
	}
	return str[0:2]
}