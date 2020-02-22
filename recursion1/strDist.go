package recursion1

import "strings"

func strDist(str string, sub string) int {
	nt := len(str)
	nu := len(sub)

	if nt < nu {
		return 0
	}

	if strings.Compare(str[0:nu], sub) != 0 {
		return strDist(str[1:nt], sub)
	}

	if strings.Compare(str[nt-nu:nt], sub) != 0 {
		return strDist(str[0:nt-1], sub)
	}

	return len(str)
}
