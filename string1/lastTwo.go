package string1

func lastTwo(str string) string {
	n := len(str)
	if n < 2 {
		return str
	}
	return str[0:n-2] + str[n-1:n] + str[n-2:n-1]
}