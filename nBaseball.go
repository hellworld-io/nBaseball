/**
    file name : nBaseball.go
 */
package main

import (
	"math/rand"
	"fmt"
	"flag"
	"time"
	"log"
	"os"
)

var arrComputerNumbers []int
var iNumberMaxLength = flag.Int("maxL", 3, "must create 3 to 5 number length")
var strCreateNumberOption = flag.String("lOpt", "1", "Option \n 1. Only first number never use 0 \n 2. All number never use 0 \n 3. 0 to 9 numbers use")

/*
	Function	: makeNumber
	Arguments	:
		iLength		= length of making numbers
		strOption
			1. Only first number never use 0
			2. All number never use 0
			3. 0 to 9 numbers use
	Desc	: to make random number
 */
func makeNumbers(iLength int, strOption string) bool{
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)
	bResult := false
	if(iLength != 0) {
		if(strOption == "1"){
			if(len(arrComputerNumbers) == 0){
				for(randomNumber == 0){
					randomNumber = rand.Intn(10)
				}
			}
		}else if(strOption == "2"){

		}else{

		}

		arrComputerNumbers = append(arrComputerNumbers, randomNumber)
		return makeNumbers(iLength-1,strOption)
	}

	if (compareNumbers()){
		bResult = true
	}else{
		bResult = false
	}

	return bResult
}

/*
	Function	: compareNumbers
	Arguments	: nothing
	Desc		: to comapare all number of arrComputerNumbers and to change duplication number
 */
func compareNumbers() bool{
	bCheck := true
	rand.Seed(int64(time.Now().Nanosecond()))
	randomNumber := rand.Intn(10)
	for i:=0; i <len(arrComputerNumbers); i++ {
		for j:=i+1; j <len(arrComputerNumbers); j++{
			fmt.Println(arrComputerNumbers[i])
			fmt.Println(arrComputerNumbers[j])
			if(arrComputerNumbers[i] == arrComputerNumbers[j]){
				for(arrComputerNumbers[i] == randomNumber || (i == 0 && randomNumber == 0)){
					randomNumber = rand.Intn(10)
				}

				arrComputerNumbers[i] = randomNumber
				bCheck = false
			}
		}
	}

	//fmt.Println("arrComputerNumbers",arrComputerNumbers)

	if(!bCheck){
		compareNumbers()
	}
	return true
}
func main() {
	flag.Parse()

	flag.Args()

	//fmt.Println(flag.Args())

	if (makeNumbers(*iNumberMaxLength, *strCreateNumberOption)){

	}else{
		log.Fatal("Making number is fail.")
		os.Exit(-1)
	}


	fmt.Println(arrComputerNumbers)
}
