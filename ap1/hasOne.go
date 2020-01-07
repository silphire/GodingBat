package ap1

func hasOne(n int) bool {
	for n != 0 {
		if n%10 == 1 {
			return true
		}
		n /= 10
	}
	return false
}
