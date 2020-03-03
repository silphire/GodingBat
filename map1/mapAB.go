package map1

func mapAB(amap map[string]string) map[string]string {
	a, okA := amap["a"]
	b, okB := amap["b"]
	if okA && okB {
		amap["ab"] = a + b
	}
	return amap
}
