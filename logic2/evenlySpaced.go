package logic2

func evenlySpaced(a int, b int, c int) bool {
	if a > b {
		t := a
		a = b
		b = t
	}
	if b > c {
		t := b
		b = c
		c = t
	}
	if a > b {
		t := a
		a = b
		b = t
	}
	return c-b == b-a
}
