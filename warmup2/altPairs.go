package warmup2

// https://codingbat.com/prob/p121596

func altPairs(str string) string {
	result := ""
	if len(str) == 1 {
		result += str[0:1]
	} else if len(str) >= 2 {
		result += str[0:2]
	}
	if len(str) == 5 {
		result += str[4:5]
	} else if len(str) >= 6 {
		result += str[4:6]
	}
	if len(str) == 9 {
		result += str[8:9]
	} else if len(str) >= 10 {
		result += str[8:10]
	}
	return result
}
