package ap1

func copyEvens(nums []int, count int) []int {
	r := make([]int, count, count)
	i := 0
	for _, n := range nums {
		if n%2 == 0 {
			r[i] = n
			i += 1
			if i == count {
				break
			}
		}
	}
	return r
}
