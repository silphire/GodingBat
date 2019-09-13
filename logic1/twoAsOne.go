package logic1

func twoAsOne(a int, b int, c int) bool {
	return a + b == c || b + c == a || c + a == b
}