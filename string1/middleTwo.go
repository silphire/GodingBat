package string1

func middleTwo(str string) string {
	n := len(str)
	return str[n/2-1:n/2+1]
}