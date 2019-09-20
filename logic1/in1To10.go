package logic1

func in1To10(n int, outsideMode bool) bool {
	result := n >= 1 && n <= 10
	if outsideMode {
		return !result
	} else {
		return result
	}
}