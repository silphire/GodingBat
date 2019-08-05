package string1

func deFront(str string) string {
	n := len(str)
	if n == 0 {
		return str
	} else if n == 1 {
		if str[0] == 'a' {
			return str
		}
		return ""
	}

	first := "a"
	if str[0] != 'a' {
		first = ""
	}

	second := "b"
	if str[1] != 'b' {
		second = ""
	}

	return first + second + str[2:n]
}