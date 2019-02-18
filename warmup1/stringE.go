package warmup1

import "strings"

func stringE(str string) bool {
	n := strings.Count(str, "e")
	return n >= 1 && n <= 3
}
