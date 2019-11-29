package array2

import "strconv"

func fizzBuzz(start int, end int) []string {
	r := make([]string, end-start, end-start)
	j := 0
	for i := start; i < end; i++ {
		if i%3 == 0 && i%5 == 0 {
			r[j] = "FizzBuzz"
		} else if i%3 == 0 {
			r[j] = "Fizz"
		} else if i%5 == 0 {
			r[j] = "Buzz"
		} else {
			r[j] = strconv.Itoa(i)
		}
		j++
	}
	return r
}
