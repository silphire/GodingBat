package warmup2

func array123(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	for i := 0; i < len(arr)-2; i++ {
		if arr[i] == 1 && arr[i+1] == 2 && arr[i+2] == 3 {
			return true
		}
	}
	return false
}
