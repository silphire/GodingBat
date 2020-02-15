package recursion1

func countPairs(str string) int {
	if len(str) < 3 {
		return 0
	}
	x := 0
	if str[0] == str[2] {
		x = 1
	}
	return x + countPairs(str[1:len(str)])
}
