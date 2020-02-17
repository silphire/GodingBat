package recursion1

func stringClean(str string) string {
	n := len(str)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return str[0:1]
	}
	if str[0] == str[1] {
		return stringClean(str[1:n])
	}
	return str[0:1] + stringClean(str[1:n])
}
