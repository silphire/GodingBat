package logic1

func less20(n int) bool {
	return (n+1)%20 == 0 || (n+2)%20 == 0
}
