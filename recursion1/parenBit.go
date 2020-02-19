package recursion1

func parenBit(str string) string {
	n := len(str)
	if n == 0 {
		return ""
	}
	if str[0] == '(' && str[n-1] == ')' {
		return str
	}

	l := 0
	if str[l] != '(' {
		l++
	}

	r := n - 1
	if str[r] != ')' {
		r--
	}

	return parenBit(str[l : r+1])
}
