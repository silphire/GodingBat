package warmup1

// https://codingbat.com/prob/p199720

func startOz(str string) string {
	if len(str) >= 2 {
		if str[0] == 'o' && str[1] == 'z' {
			return "oz"
		}
		if str[1] == 'z' {
			return "z"
		}
	}
	if len(str) > 0 && str[0] == 'o' {
		return "o"
	}
	return ""
}
