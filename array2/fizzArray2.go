package array2

import "strconv"

func fizzArray2(n int) []string {
	r := make([]string, n, n)
	for i := 0; i < n; i++ {
		r[i] = strconv.Itoa(i)
	}
	return r
}
