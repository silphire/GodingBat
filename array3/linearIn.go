package array3

func linearIn(outer []int, inner []int) bool {
	out := true
	j := 0
	for i := 0; i < len(outer); i++ {
		if out {
			if outer[i] == inner[0] {
				out = false
				j = 1
			}
		} else {
			if outer[i] == inner[j] {
				j++
				if j == len(inner) {
					return true
				}
			} else {
				out = true
			}
		}
	}
	return false
}
