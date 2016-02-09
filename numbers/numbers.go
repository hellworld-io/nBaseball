package numbers

import (
	"math/rand"
	"time"
)

var arrNumbers []int
/*
	1. Function		: MakeRandomNumbersbyComputer
	2. Arguments
		iLength		= length of making numbers
		strOption
			1. First number does not use 0 only.
			2. All numbers do not use 0.
			3. 0 to 9 numbers use
	3. Desc			: to make random number
 */
func MakeRandomNumbersByComputer(iLength int, strOption string) []int{
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)

	if(iLength != 0) {
		if(strOption == "1"){
			if(len(arrNumbers) == 0){
				for(randomNumber == 0){
					randomNumber = rand.Intn(10)
				}
			}
		}else if(strOption == "2"){
			for(randomNumber == 0){
				randomNumber = rand.Intn(10)
			}
		}

		arrNumbers = append(arrNumbers, randomNumber)
		return MakeRandomNumbersByComputer(iLength-1,strOption)
	}

	if (!compareNumbers(arrNumbers)){
		return nil
	}

	return arrNumbers
}

/*
	1. Function		: compareNumbers
	2. Arguments	: arrRandomComNumber	= random number array by Computer
	3. Desc			: to comapare all number of arrComputerNumbers and to change duplication number
 */
func compareNumbers(arrNumbers []int) bool{
	bCheck := true
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)
	for i:=0; i <len(arrNumbers); i++ {
		for j:=i+1; j <len(arrNumbers); j++{
			if(arrNumbers[i] == arrNumbers[j]){
				for(arrNumbers[i] == randomNumber || (i == 0 && randomNumber == 0)){
					randomNumber = rand.Intn(10)
				}

				arrNumbers[i] = randomNumber
				bCheck = false
			}
		}
	}

	if(!bCheck){
		compareNumbers(arrNumbers)
	}
	return true
}

func CheckUserNumber(iUserInputNumber int, arrRandomNumbers []int) string{

	return ""
}
