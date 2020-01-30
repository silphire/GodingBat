package recursion1

func sumDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumDigits(n/10)
}
