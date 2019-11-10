package string3

func sumNumbers(str string) int {
	n := 0
	d := 0
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			d = d*10 + (int)(str[i]-'0')
		} else {
			n += d
			d = 0
		}
	}
	return n + d
}
