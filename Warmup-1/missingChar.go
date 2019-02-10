package main

func missingChar(str string, n int) string {
	return str[:n] + str[n+1:]
}
