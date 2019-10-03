package logic1

func answerCell(isMorning bool, isMom bool, isAsleep bool) bool {
	if isMorning {
		return isMom
	} else if isAsleep {
		return false
	}
	return true
}
