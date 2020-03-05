package map1

func topping2(amap map[string]string) map[string]string {
	v, ok := amap["ice cream"]
	if ok {
		amap["yogurt"] = v
	}

	_, ok = amap["spinach"]
	if ok {
		amap["spinach"] = "nuts"
	}

	return amap
}
