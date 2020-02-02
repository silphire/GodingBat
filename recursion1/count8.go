package recursion1

func count8(n int) int {
	if n == 0 {
		return 0
	}
	x := 0
	if n%100 == 88 {
		x++
	}
	if n%10 == 8 {
		x++
	}
	return x + count8(n/10)
}
