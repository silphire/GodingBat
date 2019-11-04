package string2

func repeatSeparator(word string, sep string, count int) string {
	r := string(word)
	for i := 1; i < count; i++ {
		r += sep
		r += word
	}
	return r
}
