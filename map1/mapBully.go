package map1

func mapBully(amap map[string]string) map[string]string {
	val, ok := amap["a"]
	if ok {
		amap["b"] = val
		amap["a"] = ""
	}
	return amap
}
