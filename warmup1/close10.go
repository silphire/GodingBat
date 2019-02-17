package warmup1

import "math"

func close10(a int, b int) int {
	absA := math.Abs(10.0 - float64(a))
	absB := math.Abs(10.0 - float64(b))
	if absA < absB {
		return a
	} else if absA > absB {
		return b
	}
	return 0
}
