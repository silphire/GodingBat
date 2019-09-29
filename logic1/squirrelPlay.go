package logic1

func squirrelPlay(temp int, isSummer bool) bool {
	if isSummer {
		return temp >= 60 && temp <= 100
	}
	return temp >= 60 && temp <= 90
}