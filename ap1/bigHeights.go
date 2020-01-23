package ap1

func bigHeights(heights []int, start int, end int) int {
	r := 0
	for i := start; i < end; i++ {
		alt := heights[i] - heights[i+1]
		if alt <= -5 || alt >= 5 {
			r++
		}
	}

	return r
}
