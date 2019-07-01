package warmup2

func stringSplosion(str string) string {
	s := ""
	for i := 0; i < len(str); i++ {
		s += str[0 : i+1]
	}
	return s
}
