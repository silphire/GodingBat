package ap1

func wordsCount(words []string, length int) int {
	count := 0
	for _, w := range words {
		if len(w) == length {
			count++
		}
	}
	return count
}
