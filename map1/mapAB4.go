package map1

func mapAB4(amap map[string]string) map[string]string {
	va, okA := amap["a"]
	vb, okB := amap["b"]
	if okA && okB {
		if len(va) < len(vb) {
			amap["c"] = vb
		} else if len(va) > len(vb) {
			amap["c"] = va
		} else {
			amap["a"] = ""
			amap["b"] = ""
		}
	}
	return amap
}
