package string2

func oneTwo(str string) string {
	r := ""
	for i := 0; i < len(str); i += 3 {
		if len(str)-i < 3 {
			r += string(str[i:len(str)])
		} else {
			r += string(str[i+1:i+3]) + string(str[i])
		}
	}
	return r
}
