package recursion1

func changeXY(str string) string {
	if len(str) == 0 {
		return ""
	}
	if str[0] == 'x' {
		return "y" + changeXY(str[1:])
	}
	return str[0:1] + changeXY(str[1:])
}
