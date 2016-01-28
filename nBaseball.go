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
var iNumberMaxLength = flag.Int("maxL", 3, "")

/*
	Option
		1. First number never use 0
		2. All number never use 0
		3. 0 to 9 numbers use
 */
func makeNumbers(iLength int) bool{
	rand.Seed(int64(time.Now().Nanosecond()))

	if(iLength != 0) {
		randomNumber := rand.Intn(10)
		//First number never use 0 (It is option)
		if(len(arrComputerNumbers) == 0){
			for(randomNumber == 0){
				randomNumber = rand.Intn(10)
			}
		}

		arrComputerNumbers = append(arrComputerNumbers, randomNumber)
		return makeNumbers(iLength-1)
	}

	return true

}

func main() {
	flag.Parse()

	flag.Args()

	//fmt.Println(flag.Args())
	//fmt.Println(*iNumberMaxLength)


	/*
	rand.Seed(int64(time.Now().Nanosecond()))
	for i:=0; i< 10; i++{
		fmt.Println(rand.Intn(10))
	}
	*/

	if (makeNumbers(*iNumberMaxLength)){

	}

	fmt.Println(arrComputerNumbers)
}
