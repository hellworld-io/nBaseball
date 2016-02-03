package util

func CheckArgNumberLength(iNumberLength int) bool{
	if(iNumberLength < 3 || iNumberLength > 5){
		return false
	}
	return true
}