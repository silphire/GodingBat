package map2

func wordAppend(strings []string) string {
	c := map[string]int{}

	r := ""
	for _, s := range strings {
		v, ok := c[s]
		if ok {
			if v == 1 {
				c[s] = 0
				r += s
			} else {
				c[s]++
			}
		} else {
			c[s] = 1
		}
	}
	return r
}
