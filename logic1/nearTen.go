package logic1

func nearTen(num int) bool {
	n := num % 10
	return n >= 8 || n <= 2
}