package warmup1

func everyNth(str string, n int) string {
	size := len(str) / n
	if len(str)%n > 0 {
		size++
	}

	var result = make([]byte, size)
	for i := 0; i < size; i++ {
		result[i] = str[i*n]
	}
	return string(result[:])
}
