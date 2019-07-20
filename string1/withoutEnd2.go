package string1

func withoutEnd2(str string) string {
	if len(str) <= 2 {
		return ""
	}
	return str[1:len(str)-1]
}