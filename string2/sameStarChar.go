package string2

func sameStarChar(str string) bool {
	ln := len(str)
	for i := 0; i < ln; i++ {
		if str[i] == '*' {
			if i == 0 || i == ln-1 {
				continue
			}
			if str[i-1] != str[i+1] {
				return false
			}
		}
	}
	return true
}
