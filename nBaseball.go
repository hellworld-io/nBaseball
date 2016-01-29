/**
    file name : nBaseball.go
 */
package main

import (
	"math/rand"
	"fmt"
	"flag"
	"time"
)

var arrComputerNumbers []int
var iNumberMaxLength = flag.Int("maxL", 3, "must create 3 to 5 number length")
var strCreateNumberOption = flag.String("lOpt", "1", "Option \n 1. Only first number never use 0 \n 2. All number never use 0 \n 3. 0 to 9 numbers use")

/*
	Option
		1. Only first number never use 0
		2. All number never use 0
		3. 0 to 9 numbers use
 */
func makeNumbers(iLength int, strOption string) bool{
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)

	if(iLength != 0) {
		if(strOption == "1"){
			if(len(arrComputerNumbers) == 0){
				for(randomNumber == 0){
					randomNumber = rand.Intn(10)
				}
			}else{
				for v := range arrComputerNumbers {
					if(arrComputerNumbers[v] == randomNumber){
						randomNumber = rand.Intn(10)
					}
				}
			}
		}else if(strOption == "2"){

		}else{

		}

		arrComputerNumbers = append(arrComputerNumbers, randomNumber)
		return makeNumbers(iLength-1,strOption)
	}else{
		for i:=0; i< len(arrComputerNumbers)-1 ; i++ {
			if(arrComputerNumbers[*iNumberMaxLength-1] == arrComputerNumbers[i]){
				fmt.Println("There is duplication of number")
				fmt.Println("arrComputerNumbers[*iNumberMaxLength-1]",arrComputerNumbers[*iNumberMaxLength-1])
				fmt.Println("arrComputerNumbers[i]",arrComputerNumbers[i])
				randomNumber = rand.Intn(10)
				fmt.Println("randomNumber",randomNumber)
				if(randomNumber == arrComputerNumbers[*iNumberMaxLength-1]){
					randomNumber = rand.Intn(10)
				}
				arrComputerNumbers[i] = randomNumber
			}
		}
	}

	return true

}

func main() {
	flag.Parse()

	flag.Args()

	fmt.Println(flag.Args())

	if (makeNumbers(*iNumberMaxLength, *strCreateNumberOption)){

	}


	fmt.Println(arrComputerNumbers)
}
