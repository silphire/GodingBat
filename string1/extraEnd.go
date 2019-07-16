package string1

func extraEnd(str string) string {
	n := len(str)
	s := str[n-2:n]
	return s + s + s
}