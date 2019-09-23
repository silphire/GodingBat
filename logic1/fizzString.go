package logic1

import "strings"

func fizzString(str string) string {
	f := strings.HasPrefix(str, "f")
	b := strings.HasSuffix(str, "b")

	if f && b {
		return "FizzBuzz"
	}
	if f {
		return "Fizz"
	}
	if b {
		return "Buzz"
	}
	return str
}
