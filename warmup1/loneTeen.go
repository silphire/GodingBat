package warmup1

// https://codingbat.com/prob/p165701

func loneTeen(a int, b int) bool {
	f1 := a >= 13 && a <= 19
	f2 := b >= 13 && b <= 19
	return f1 && !f2 || !f1 && f2
}
