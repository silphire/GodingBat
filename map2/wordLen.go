package map2

func wordLen(strings []string) map[string]int {
	r := map[string]int{}

	for _, s := range strings {
		r[s] = len(s)
	}

	return r
}
