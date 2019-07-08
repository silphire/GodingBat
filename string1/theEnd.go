package string1

func theEnd(str string, front bool) string {
	if front {
		return str[0:1]
	} else {
		return str[len(str)-1 : len(str)]
	}
}
