package string3

func maxBlock(str string) int {
	var p rune
	r := 0
	n := 0

	for _, c := range str {
		if p == c {
			n++
		} else {
			if r < n {
				r = n
			}
			n = 1
			p = c
		}
	}

	if r < n {
		return n
	}
	return r
}
