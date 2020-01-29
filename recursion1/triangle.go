package recursion1

func triangle(rows int) int {
	if rows == 0 {
		return 0
	}
	return rows + triangle(rows-1)
}
