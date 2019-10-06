package logic1

func greenTicket(a int, b int, c int) int {
	if a == b && b == c {
		return 20
	}
	if a == b || b == c || c == a {
		return 10
	}
	return 0
}
