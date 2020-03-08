package map1

import "strings"

func mapAB2(amap map[string]string) map[string]string {
	va, okA := amap["a"]
	vb, okB := amap["b"]
	if okA && okB && strings.Compare(va, vb) == 0 {
		delete(amap, "a")
		delete(amap, "b")
	}
	return amap
}
