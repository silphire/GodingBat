package logic1

func inOrderEqual(a int, b int, c int, equalOk bool) bool {
	if equalOk {
		return a <= b && b <= c
	}
	return a < b && b < c
}
