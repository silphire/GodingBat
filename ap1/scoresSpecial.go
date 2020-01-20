package ap1

func scoresSpecial(a []int, b []int) int {
	ma := 0
	mb := 0
	for _, x := range a {
		if x%10 == 0 && ma < x {
			ma = x
		}
	}
	for _, x := range b {
		if x%10 == 0 && mb < x {
			mb = x
		}
	}
	return ma + mb
}
