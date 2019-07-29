package string1

func firstTwo(str string) string {
	n := len(str)
	if n > 2 {
		n = 2
	}
	return str[0:n]
}