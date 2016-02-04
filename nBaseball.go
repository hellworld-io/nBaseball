/**
    file name : nBaseball.go
 */
package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	_"bufio"
	"nBaseball/util"
	"nBaseball/numbers"
)

var arrComputerNumbers []int
var iNumberMaxLength = flag.Int("maxL", 3, "must create 3 to 5 number length")
var strCreateNumberOption = flag.String("lOpt", "1", "Option \n 1. First number does no use 0 only. \n 2. All number do not use 0. \n 3. 0 to 9 numbers use")




func main() {
	flag.Parse()

	flag.Args()

	//fmt.Println(flag.Args())
	//fmt.Println("*iNumberMaxLength",*iNumberMaxLength)


	if(!util.CheckArgNumberLength(*iNumberMaxLength)){
		log.Fatal("Making number length is 3 to 5.")
		os.Exit(-1)
	}

	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	if line == "exit" {
	//		os.Exit(0)
	//	}
	//	fmt.Println(line) // Println will add back the final '\n'
	//}
	//if err := scanner.Err(); err != nil {
	//	fmt.Fprintln(os.Stderr, "reading standard input:", err)
	//}

	arrComputerNumbers = numbers.MakeRandomNumbersbyComputer(*iNumberMaxLength, *strCreateNumberOption)
	if (arrComputerNumbers != nil){

	}else{
		log.Fatal("Making number is fail.")
		os.Exit(-1)
	}


	fmt.Println(arrComputerNumbers)
}
