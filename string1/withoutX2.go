package string1

func withoutX2(str string) string {
	result := ""
	if str[0] != 'x' {
		result += str[0:1]
	}
	if str[1] != 'x' {
		result += str[1:2]
	}
	return result + str[2:len(str)]
}