package recursion1

func count7(n int) int {
	if n == 0 {
		return 0
	}
	x := 0
	if n%10 == 7 {
		x = 1
	}
	return x + count7(n/10)
}
