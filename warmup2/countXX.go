package warmup2

func countXX(str string) int {
	n := 0
	for i := 0; i < len(str)-1; i++ {
		if str[i] == 'x' && str[i+1] == 'x' {
			n++
		}
	}
	return n
}
