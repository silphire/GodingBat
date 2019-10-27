package string2

func repeatFront(str string, n int) string {
	r := ""
	for i := n; i > 0; i-- {
		r += string(str[0:i])
	}
	return r
}
