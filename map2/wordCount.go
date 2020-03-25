package map2

func wordCount(strings []string) map[string]int {
	amap := map[string]int{}

	for _, s := range strings {
		_, ok := amap[s]
		if !ok {
			amap[s] = 0
		}
		amap[s]++
	}

	return amap
}
