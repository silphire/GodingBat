package ap1

func wordsWithoutList(words []string, length int) []string {
	n := 0
	for i := 0; i < len(words); i++ {
		if len(words[i]) != length {
			n++
		}
	}

	j := 0
	r := make([]string, n, n)
	for i := 0; i < len(words); i++ {
		if len(words[i]) != length {
			r[j] = words[i]
			j++
		}
	}

	return r
}
