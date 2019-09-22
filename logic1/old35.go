package logic1

func old35(n int) bool {
	return n%3 == 0 && n%5 != 0 || n%3 != 0 && n%5 == 0
}
