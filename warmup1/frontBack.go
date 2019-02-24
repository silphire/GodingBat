package warmup1

// https://codingbat.com/prob/p123384

func frontBack(str string) string {
	n := len(str)
	if n == 1 {
		return str
	}
	return str[n-1:n] + str[1:n-1] + str[0:1]
}
