package string3

func mirrorEnds(str string) string {
	i := 0
	for ; i < len(str); i++ {
		if str[i] != str[len(str)-1-i] {
			break
		}
	}
	return string(str[0:i])
}
