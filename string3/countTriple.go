package string3

func countTriple(str string) int {
	n := 0
	for i := 0; i < len(str)-2; i++ {
		f := true
		c := str[i]
		for j := 0; j < 3; j++ {
			if str[i+j] != c {
				f = false
				break
			}
		}
		if f {
			n++
		}
	}
	return n
}
