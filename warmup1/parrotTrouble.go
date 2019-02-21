package warmup1

func parrotTrouble(talking bool, hour int) bool {
	return talking && (hour < 7 || hour > 20)
}
