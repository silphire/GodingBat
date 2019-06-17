// https://codingbat.com/prob/p142270

package warmup2

func stringTimes(str string, n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret += str
	}
	return ret
}
