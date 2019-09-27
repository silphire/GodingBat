package logic1

func redTicket(a int, b int, c int) int {
	if b == c {
		if a == b {
			if a == 2 {
				return 10
			}
			return 5
		}
		return 1
	}
	return 0
}
