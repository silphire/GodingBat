package recursion1

func countX(str string) int {
	if len(str) == 0 {
		return 0
	}
	r := 0
	if str[0] == 'x' {
		r = 1
	}
	return r + countX(str[1:len(str)])
}
