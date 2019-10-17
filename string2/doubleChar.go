package string2

func doubleChar(str string) string {
	s := ""
	for _, c := range str {
		ch := string(c)
		s += ch + ch
	}
	return s
}
