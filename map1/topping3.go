package map1

func topping3(amap map[string]string) map[string]string {
	v, ok := amap["potato"]
	if ok {
		amap["fries"] = v
	}

	v, ok = amap["salad"]
	if ok {
		amap["spinach"] = v
	}

	return amap
}
