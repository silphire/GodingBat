package map1

func mapAB3(amap map[string]string) map[string]string {
	va, okA := amap["a"]
	vb, okB := amap["b"]
	if okA && okB {
		return amap
	}
	if okA {
		amap["b"] = va
	}
	if okB {
		amap["a"] = vb
	}
	return amap
}
