package warmup2

func stringX(str string) string {
	r := ""
	for i := 0; i < len(str); i++ {
		if i == 0 || i == len(str)-1 || str[i] != 'x' {
			r += str[i : i+1]
		}
	}
	return r
}
