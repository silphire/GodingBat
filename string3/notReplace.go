package string3

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func notReplace(str string) string {
	is := "is"
	r := ""
	for i := 0; i < len(str); i++ {
		if i > len(str)-len(is) {
			r += string(str[i])
			continue
		}

		f := true
		for j := 0; j < len(is); j++ {
			if str[i+j] != is[j] {
				f = false
				break
			}
		}
		if f {
			if (i == 0 || !isLetter(str[i-1])) && (i == len(str)-len(is) || !isLetter(str[i+len(is)])) {
				r += "is not"
				i++
				continue
			}
		}
		r += string(str[i])
	}
	return r
}
