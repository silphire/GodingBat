package ap1

func dividesSelf(n int) bool {
	nn := n
	for n > 0 {
		x := n % 10
		if x == 0 {
			return false
		}
		if nn%x != 0 {
			return false
		}
		n /= 10
	}
	return true
}
