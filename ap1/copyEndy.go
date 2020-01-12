package ap1

func copyEndy(nums []int, count int) []int {
	r := make([]int, count, count)
	i := 0
	for _, n := range nums {
		if n >= 0 && n <= 10 || n >= 90 && n <= 100 {
			r[i] = n
			i++
			if i == count {
				break
			}
		}
	}
	return r
}
