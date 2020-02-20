package recursion1

func nestParen(str string) bool {
	if len(str) == 0 {
		return true
	}
	if len(str) < 2 {
		return false
	}
	if str[0] == '(' && str[len(str)-1] == ')' {
		return nestParen(str[1 : len(str)-1])
	}
	return false
}
