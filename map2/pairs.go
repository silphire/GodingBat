package map2

func pairs(strings []string) map[string]string {
	r := map[string]string{}
	for _, s := range strings {
		n := len(s)
		r[s[0:1]] = string(s[n-1 : n])
	}
	return r
}
