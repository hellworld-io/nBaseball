package numbers

import (
	"math/rand"
	"time"
)

var randomNumbers []int
/*
	1. Function		: MakeRandomNumbers
	2. Arguments
		length		: length of making numbers
		option
			1. First number does not use 0 only.
			2. All numbers do not use 0.
			3. 0 to 9 numbers use
	3. Desc			: to make random number
	4. Return		: []int
 */
func MakeRandomNumbers(length int, option string) []int{
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)

	if length != 0 {
		if option == "1" {
			for randomNumber == 0 {
				randomNumber = rand.Intn(10)
			}
		}else if option == "2" {
			if len(randomNumbers) == 0 {
				for randomNumber == 0 {
					randomNumber = rand.Intn(10)
				}
			}
		}

		randomNumbers = append(randomNumbers, randomNumber)
		return MakeRandomNumbers(length-1,option)
	}

	if  !removeDupleNumbers(randomNumbers) {
		return randomNumbers
	}

	return randomNumbers
}

/*
	1. Function		: removeDupleNumbers
	2. Arguments		: the numbers is random int slice.
	3. Desc			: to change duplication number
	4. Return		: bool
 */
func removeDupleNumbers(numbers []int) bool{
	bCheck := true
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)

	for i:=0; i <len(numbers); i++ {
		for j:=i+1; j <len(numbers); j++{
			if numbers[i] == numbers[j] {
				for numbers[i] == randomNumber || (i == 0 && randomNumber == 0) {
					randomNumber = rand.Intn(10)
				}

				numbers[i] = randomNumber
				bCheck = false
			}
		}
	}

	if !bCheck {
		removeDupleNumbers(numbers)
	}
	return true
}