package string2

func countCode(str string) int {
	target := "code"
	n := 0
	for i := 0; i < len(str)-4+1; i++ {
		isFound := true
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue
			}
			if str[i+j] != target[j] {
				isFound = false
				break
			}
		}
		if isFound {
			n++
		}
	}
	return n
}
