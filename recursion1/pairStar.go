package recursion1

func pairStar(str string) string {
	if len(str) <= 1 {
		return str
	}
	if str[0] == str[1] {
		return str[0:1] + "*" + pairStar(str[1:len(str)])
	}
	return str[0:1] + pairStar(str[1:len(str)])
}
