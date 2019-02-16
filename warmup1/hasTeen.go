package warmup1

func hasTeen(nums ...int) bool {
	for _, num := range nums {
		if num >= 13 && num <= 19 {
			return true
		}
	}
	return false
}
