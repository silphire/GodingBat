package string1

func nTwice(str string, n int) string {
	return str[0:n] + str[len(str)-n:len(str)]
}