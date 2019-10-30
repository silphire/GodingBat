package string2

func plusOut(str string, word string) string {
	r := ""
	for i := 0; i < len(str); i++ {
		f := true
		for j := 0; j < len(word); j++ {
			if str[i] == word[j] {
				f = false
				break
			}
		}
		if f {
			r += "+"
		} else {
			r += string(str[i])
		}
	}
	return r
}
