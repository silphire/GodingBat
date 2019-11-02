package string2

import "regexp"

func xyzThere(str string) bool {
	r := regexp.MustCompile(`^xyz|[^.]xyz`)
	return r.MatchString(str)
}
