package ap1

func sumHeights(heights []int, start int, end int) int {
	r := 0
	for i := start; i < end; i++ {
		alt := heights[i] - heights[i+1]
		if alt < 0 {
			alt = -alt
		}
		r += alt
	}

	return r
}
