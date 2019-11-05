package string2

func zipZap(str string) string {
	r := ""
	for i := 0; i < len(str); i++ {
		if i+3 <= len(str) && str[i] == 'z' && str[i+2] == 'p' {
			r += "zp"
			i += 2
		} else {
			r += string(str[i])
		}
	}
	return r
}
