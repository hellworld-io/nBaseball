package util

/*
	1. Function		: CheckUserNumberLength
	2. Arguments	:
		strUserNumberLength	= user input numbers
		iMaxNumberLength	= random number max length
	3. Desc			: to compare max number length with user input numbers
 */
func CheckUserNumberLength(strUserInputNumber string, iMaxNumberLength int) bool{
	if len(strUserInputNumber) != iMaxNumberLength {
		return false
	}

	return true
}

/*
	1. Function		: CheckUserInputNumber
	2. Arguments	:
		strUsrInputNumbers	= user input numbers
	3. Desc			: to check duplication numbers for inputted user numbers
 */
func CheckUserInputNumber(strUsrInputNumbers string) bool{
	for idx := 0; idx < len(strUsrInputNumbers); idx++{
		for jdx := idx+1; jdx < len(strUsrInputNumbers); jdx++{
			if strUsrInputNumbers[idx] == strUsrInputNumbers[jdx]{
				return false
			}
		}
	}
	return true
}