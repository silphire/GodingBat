package warmup2

func stringMatch(a string, b string) int {
	if len(a) <= 1 || len(b) <= 1 {
		return 0
	}

	sz := 0
	if len(a) < len(b) {
		sz = len(a)
	} else {
		sz = len(b)
	}

	n := 0
	for i := 0; i < sz-1; i++ {
		if a[i] == b[i] && a[i+1] == b[i+1] {
			n++
		}
	}
	return n
}
