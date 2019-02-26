package warmup1

func icyHot(temp1 int, temp2 int) bool {
	return temp1 < 0 && temp2 > 100 || temp1 > 100 && temp2 < 0
}
