package recursion1

func allStar(str string) string {
	if len(str) == 1 {
		return str
	}
	return str[0:1] + "*" + allStar(str[1:len(str)])
}
