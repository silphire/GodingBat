package recursion1

import "strings"

func changePi(str string) string {
	if len(str) < 2 {
		return str
	}
	if strings.Compare(str[0:2], "pi") == 0 {
		return "3.14" + changePi(str[2:len(str)])
	}
	return str[0:1] + changePi(str[1:len(str)])
}
