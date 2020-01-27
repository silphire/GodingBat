package recursion1

func bunnyEars(bunnies int) int {
	if bunnies == 0 {
		return 0
	}
	if bunnies == 1 {
		return 2
	}
	return 2*bunnyEars(bunnies/2) + 2*(bunnies%2)
}
