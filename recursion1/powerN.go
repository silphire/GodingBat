package recursion1

func powerN(base int, n int) int {
	if n == 0 {
		return 1
	}
	return base * powerN(base, n-1)
}
