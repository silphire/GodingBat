package string2

func starOut(str string) string {
	r := ""
	for i := 0; i < len(str); i++ {
		if str[i] == '*' {
			continue
		}
		if i > 0 && str[i-1] == '*' {
			continue
		}
		if i < len(str)-1 && str[i+1] == '*' {
			continue
		}
		r += string(str[i])
	}
	return r
}
