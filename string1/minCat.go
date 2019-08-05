package string1

func minCat(a string, b string) string {
	na := len(a)
	nb := len(b)
	if na < nb {
		return a + b[nb-na:nb]
	} else if na > nb {
		return a[na-nb:na] + b
	} else {
		return a + b
	}
}