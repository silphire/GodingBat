package warmup1

func missingChar(str string, n int) string {
	return str[:n] + str[n+1:]
}
