package logic1

func inOrder(a int, b int, c int, bOk bool) bool {
	if bOk {
		return c > b
	}
	return b > a && c > b
}
