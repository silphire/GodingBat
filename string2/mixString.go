package string2

func mixString(a string, b string) string {
	r := ""
	la := len(a)
	lb := len(b)
	lm := la
	if lm < lb {
		lm = lb
	}
	for i := 0; i < lm; i++ {
		if i < la {
			r += string(a[i])
		}
		if i < lb {
			r += string(b[i])
		}
	}
	return r
}
