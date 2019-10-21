package string2

func repeatEnd(str string, n int) string {
	ln := len(str)
	r := ""
	for i := 0; i < n; i++ {
		r += string(str[ln-n : ln])
	}
	return r
}
