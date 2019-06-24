package warmup2

// https://codingbat.com/prob/p101475

func frontTimes(str string, n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += str[0:3]
	}
	return ret
}
