// https://codingbat.com/prob/p161642

package warmup1

func backAround(str string) string {
	result := make([]byte, len(str)+2)
	copy(result[1:len(str)+1], str[:])
	result[0] = str[len(str)-1]
	result[len(result)-1] = str[len(str)-1]
	return string(result)
}
