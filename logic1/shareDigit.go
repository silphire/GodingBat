package logic1

func shareDigit(a int, b int) bool {
	d1 := a / 10
	d2 := a % 10
	d3 := b / 10
	d4 := b % 10
	return d1 == d3 || d1 == d4 || d2 == d3 || d2 == d4
}
