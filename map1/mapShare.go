package map1

func mapShare(amap map[string]string) map[string]string {
	val, ok := amap["a"]
	if ok {
		amap["b"] = val
	}
	delete(amap, "c")

	return amap
}
