package numbers

import (
	"math/rand"
	"time"
)

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
func MakeRandomNumbersbyComputer(iLength int, strOption string, arrRandomComNumber []int) bool{
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)
	bResult := false

	if(iLength != 0) {
		if(strOption == "1"){
			if(len(arrRandomComNumber) == 0){
				for(randomNumber == 0){
					randomNumber = rand.Intn(10)
				}
			}
		}else if(strOption == "2"){

		}else{

		}

		arrRandomComNumber = append(arrRandomComNumber, randomNumber)
		return MakeRandomNumbersbyComputer(iLength-1,strOption,arrRandomComNumber)
	}

	if (compareNumbers(arrRandomComNumber)){
		bResult = true
	}else{
		bResult = false
	}

	return bResult
}

/*
	1. Function		: compareNumbers
	2. Arguments	: arrRandomComNumber	= random number array by Computer
	3. Desc			: to comapare all number of arrComputerNumbers and to change duplication number
 */
func compareNumbers(arrRandomComNumber []int) bool{
	bCheck := true
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)
	for i:=0; i <len(arrRandomComNumber); i++ {
		for j:=i+1; j <len(arrRandomComNumber); j++{
			if(arrRandomComNumber[i] == arrRandomComNumber[j]){
				for(arrRandomComNumber[i] == randomNumber || (i == 0 && randomNumber == 0)){
					randomNumber = rand.Intn(10)
				}

				arrRandomComNumber[i] = randomNumber
				bCheck = false
			}
		}
	}

	if(!bCheck){
		compareNumbers(arrRandomComNumber)
	}
	return true
}
