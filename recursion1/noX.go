package recursion1

func noX(str string) string {
	if len(str) == 0 {
		return ""
	}
	if str[0] == 'x' {
		return noX(str[1:len(str)])
	}
	return str[0:1] + noX(str[1:len(str)])
}
