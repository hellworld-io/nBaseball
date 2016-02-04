package util

/*
	1. Function		: CheckArgNumberLength
	2. Arguments	: iNumberLength	= to make max length of number
	3. Desc			: to check max length
 */
func CheckArgNumberLength(iNumberLength int) bool{
	if(iNumberLength < 3 || iNumberLength > 5){
		return false
	}
	return true
}