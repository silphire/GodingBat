package ap1

import "strings"

func userCompare(aName string, aId int, bName string, bId int) int {
	r := strings.Compare(aName, bName)
	if r != 0 {
		return r
	}
	if aId < bId {
		return -1
	}
	if aId > bId {
		return 1
	}
	return 0
}
