package string2

func xyBalance(str string) bool {
	f := false
	for i := 0; i < len(str); i++ {
		if f {
			if str[i] == 'y' {
				return true
			}
		} else {
			if str[i] == 'x' {
				f = true
			}
		}
	}
	return false
}
