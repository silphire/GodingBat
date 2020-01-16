package ap1

func matchUp(a []string, b []string) int {
	n := 0
	for i := 0; i < len(a); i++ {
		if len(a[i]) > 0 && len(b[i]) > 0 && a[i][0] == b[i][0] {
			n++
		}
	}
	return n
}
