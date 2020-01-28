package recursion1

func bunnyEars2(bunnies int) int {
	if bunnies == 0 {
		return 0
	}
	return 3 - (bunnies % 2) + bunnyEars2(bunnies-1)
}
