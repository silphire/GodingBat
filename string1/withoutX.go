package string1

func withoutX(str string) string {
	if len(str) > 0 && str[0] == 'x' {
		str = str[1:len(str)]
	}
	if len(str) > 0 && str[len(str) - 1] == 'x' {
		str = str[0:len(str) - 1]
	}
	return str
}