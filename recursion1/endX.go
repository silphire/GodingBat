package recursion1

func endX(str string) string {
	if len(str) == 0 {
		return ""
	}
	if str[0] == 'x' {
		return endX(str[1:len(str)]) + "x"
	}
	return str[0:1] + endX(str[1:len(str)])
}
