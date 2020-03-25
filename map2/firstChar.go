package map2

func firstChar(strings []string) map[string]string {
	amap := map[string]string{}

	for _, s := range strings {
		k := s[0:1]
		_, ok := amap[k]
		if !ok {
			amap[k] = ""
		}
		amap[k] = amap[k] + s
	}

	return amap
}
