package logic1

func lastDigit(a int, b int, c int) bool {
	aa := a % 10
	bb := b % 10
	cc := c % 10
	return aa == bb || bb == cc || cc == aa
}