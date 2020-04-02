package map2

func wordMultiple(strings []string) map[string]bool {
	amap := map[string]bool{}
	for _, s := range strings {
		v, ok := amap[s]
		if !ok {
			amap[s] = false
		} else if !v {
			amap[s] = true
		}
	}
	return amap
}
