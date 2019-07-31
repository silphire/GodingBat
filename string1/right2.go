package string1

func right2(str string) string {
	n := len(str)
	if n > 2 {
		return str[n-2 : n] + str[0:n-2]
	} else {
		return str
	}
}