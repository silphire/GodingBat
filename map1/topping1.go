package map1

func topping1(amap map[string]string) map[string]string {
	_, ok := amap["ice cream"]
	if ok {
		amap["ice cream"] = "cherry"
	}
	amap["bread"] = "butter"
	return amap
}
