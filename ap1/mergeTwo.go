package ap1

import "strings"

func mergeTwo(a []string, b []string, n int) []string {
	c := make([]string, n, n)
	ia := 0
	ib := 0
	for ic := 0; ic < n; ic++ {
		f := 0
		if ia >= len(a) {
			f = 1
		} else if ib >= len(b) {
			f = -1
		} else {
			f = strings.Compare(a[ia], b[ib])
		}

		if f < 0 {
			c[ic] = a[ia]
			ia++
		} else if f > 0 {
			c[ic] = b[ib]
			ib++
		} else {
			if ia < len(a) {
				c[ic] = a[ia]
			} else if ib < len(b) {
				c[ic] = b[ib]
			}
			ia++
			ib++
		}

		for ia < len(a) && strings.Compare(c[ic], a[ia]) == 0 {
			ia++
		}

		for ib < len(b) && strings.Compare(c[ic], b[ib]) == 0 {
			ib++
		}
	}

	return c
}
