/**
    file name : nBaseball.go
 */
package main

import (
	"math/rand"
	"fmt"
	"flag"
	_"time"
)

var arrComputerNumbers []int
var iNumberMaxLength = flag.Int("maxL", 3, "")

func makeNumbers(iLength int) int{
	randomNumber := rand.New(rand.NewSource(10))

	if(iLength != 0) {
		arrComputerNumbers = append(arrComputerNumbers, randomNumber.Intn(10))
		return makeNumbers(iLength-1)
	}else{
		return 0
	}


}

func main() {
	flag.Parse()

	flag.Args()

	//fmt.Println(flag.Args())
	//fmt.Println(*iNumberMaxLength)

	/*
	rand.Seed(int64(time.Now().Nanosecond()))
	for i:=0; i< 5; i++{
		fmt.Println(rand.Intn(10))
	}
	*/

	for i:=0; i< 5; i++{
		fmt.Println(rand.Intn(10))
	}

	makeNumbers(*iNumberMaxLength)

	fmt.Println(arrComputerNumbers)
}
