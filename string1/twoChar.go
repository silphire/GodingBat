package string1

func twoChar(str string, index int) string {
	if index + 2 > len(str) {
		return str[0:2]
	}
	return str[index:index+2]
}