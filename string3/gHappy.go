package string3

func gHappy(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != 'g' {
			continue
		}
		if !(i > 0 && str[i-1] == 'g' || i < len(str)-1 && str[i+1] == 'g') {
			return false
		}
	}
	return true
}
