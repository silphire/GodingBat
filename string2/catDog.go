package string2

import "strings"

func catDog(str string) bool {
	return strings.Count(str, "dog") == strings.Count(str, "cat")
}
