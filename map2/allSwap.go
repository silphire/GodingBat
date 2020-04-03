package map2

func allSwap(strings []string) []string {
	prev := map[byte]int{}
	for i := range strings {
		c := strings[i][0]
		j, ok := prev[c]
		if ok {
			strings[i], strings[j] = strings[j], strings[i]
			delete(prev, c)
		} else {
			prev[c] = i
		}
	}
	return strings
}
