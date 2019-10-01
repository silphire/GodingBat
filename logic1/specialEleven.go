package logic1

func specialEleven(n int) bool {
	x := n % 11
	return x == 0 || x == 1
}
