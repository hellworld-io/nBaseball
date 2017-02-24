package numbers

import (
	"math/rand"
	"time"
	"strconv"
)

var ArrRandomNumbers []int
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
func MakeRandomNumbersByComputer(iLength int, strOption string) bool{
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)

	if iLength != 0 {
		if strOption == "1" {
			for randomNumber == 0 {
				randomNumber = rand.Intn(10)
			}
		}else if strOption == "2" {
			if len(ArrRandomNumbers) == 0 {
				for randomNumber == 0 {
					randomNumber = rand.Intn(10)
				}
			}
		}

		ArrRandomNumbers = append(ArrRandomNumbers, randomNumber)
		return MakeRandomNumbersByComputer(iLength-1,strOption)
	}

	if  !compareRandomNumbersAndResetNumbers(ArrRandomNumbers) {
		return true
	}

	return true
}

/*
	1. Function		: compareRandomNumbersAndResetNumbers
	2. Arguments	: arrRandomComNumber	= random number array by Computer
	3. Desc			: to comapare all number of arrComputerNumbers and to change duplication number
 */
func compareRandomNumbersAndResetNumbers(arrNumbers []int) bool{
	bCheck := true
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)
	for i:=0; i <len(arrNumbers); i++ {
		for j:=i+1; j <len(arrNumbers); j++{
			if rrNumbers[i] == arrNumbers[j] {
				for arrNumbers[i] == randomNumber || (i == 0 && randomNumber == 0) {
					randomNumber = rand.Intn(10)
				}

				arrNumbers[i] = randomNumber
				bCheck = false
			}
		}
	}

	if !bCheck {
		compareRandomNumbersAndResetNumbers(arrNumbers)
	}
	return true
}


func CompareUserNumbersAndRandomNumbers(strUserInputNumbers string, arrRandomNumbers []int) string{
	var strResul string = ""
	var iStrik int = 0
	var iBall int = 0

	for idx, randomNumber := range arrRandomNumbers {
		for jdx, strUserInputNumber := range strUserInputNumbers {
			if  strconv.Itoa(randomNumber) == string(strUserInputNumber) {
				if  idx == jdx {
					iStrik++
				}else{
					iBall++
				}
			}

		}
	}
	if iStrik == 3 {
		strResul = "Out"
	}else{
		strResul = strconv.Itoa(iStrik) + "Strikes " + strconv.Itoa(iBall) + "Balls"
	}

	return strResul
}
