package map2

func word0(strs []string) map[string]int {
	r := map[string]int{}
	for _, s := range strs {
		r[s] = 0
	}
	return r
}
