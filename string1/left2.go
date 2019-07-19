package string1

func left2(str string) string {
	if len(str) <= 2 {
		return str
	}
	return str[2:len(str)] + str[0:2]
}