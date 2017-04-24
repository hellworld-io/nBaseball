package play

import (
	"strconv"
	"errors"
)

func checkInputValues(inputs string, maxLen int) error {

	if _, err := strconv.Atoi(inputs); err != nil {
		return errors.New("The input value is not number.")
	}

	if maxLen != len(inputs) {
		return errors.New("Input number of length is " + strconv.Itoa(maxLen))
	}

	return nil
}

func PlayBaseball(inputValues string, randomNumbers []int, maxLen int) (string, error) {
	var iStrik int = 0
	var iBall int = 0
	var result string

	if err := checkInputValues(inputValues, maxLen); err != nil {
		return "", err
	}

	for idx, randomNumber := range randomNumbers {
		for jdx, inputValue := range inputValues {
			if  strconv.Itoa(randomNumber) == string(inputValue) {
				if  idx == jdx {
					iStrik++
				}else{
					iBall++
				}
			}

		}
	}

	if iStrik == 3 {
		result = "Out"
	}else{
		result = strconv.Itoa(iStrik) + "Strikes " + strconv.Itoa(iBall) + "Balls"
	}

	return result, nil
}
