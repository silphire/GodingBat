package map2

func firstSwap(strings []string) []string {
	amap := map[byte]int{}
	for i, s := range strings {
		c := s[0]
		p, ok := amap[c]
		if ok {
			if p >= 0 {
				t := strings[i]
				strings[i] = strings[p]
				strings[p] = t
				amap[c] = -1
			}
		} else {
			amap[c] = i
		}
	}
	return strings
}
