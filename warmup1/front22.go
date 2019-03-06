package warmup1

// https://codingbat.com/prob/p183592

func front22(str string) string {
	if len(str) < 2 {
		return str + str + str
	}
	return str[0:2] + str + str[0:2]
}
