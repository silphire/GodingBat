package logic2

func round10(num int) int {
	if num % 10 < 5 {
		return num / 10 * 10;
	}
	return (num / 10 + 1) * 10;
}

func roundSum(a int, b int, c int) int {
	return round10(a) + round10(b) + round10(c)
}