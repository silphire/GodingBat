package logic1

func love6(a int, b int) bool {
	return a == 6 || b == 6 || a + b == 6 || a - b == 6 || b - a == 6
}