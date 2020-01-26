package ap1

import "strings"

func commonTwo(a []string, b []string) int {
	n := 0
	ia := 0
	ib := 0
	for ia < len(a) && ib < len(b) {
		f := strings.Compare(a[ia], b[ib])
		if f < 0 {
			ia++
		} else if f > 0 {
			ib++
		} else {
			n++
			ia++
			ib++
		}
	}
	return n
}
