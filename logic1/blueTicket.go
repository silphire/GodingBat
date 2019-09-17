package logic1

func blueTicket(a int, b int, c int) int {
	if a + b == 10 || b + c == 10 || c + a == 10 {
		return 10;
	}
	if a + b - (b + c) == 10 || a + b - (c + a) == 10 {
		return 5;
	}
	return 0;
}