package string1

func middleThree(str string) string {
	if len(str) == 3 {
		return str
	}
	return str[1:4]
}