package string2

func bobThere(str string) bool {
	target := "bob"
	ln := len(target)
	for i := 0; i < len(str)-ln+1; i++ {
		isFound := true
		for j := 0; j < ln; j++ {
			if j == 1 {
				continue
			}
			if str[i+j] != target[j] {
				isFound = false
				break
			}
		}
		if isFound {
			return true
		}
	}
	return false
}
