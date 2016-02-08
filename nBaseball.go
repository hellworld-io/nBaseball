/**
    file name : nBaseball.go
 */
package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"bufio"
	"nBaseball/util"
	"nBaseball/numbers"
	"time"
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

	arrComputerNumbers = numbers.MakeRandomNumbersbyComputer(*iNumberMaxLength, *strCreateNumberOption)

	if (arrComputerNumbers != nil){
		startSecond := time.Now()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Computer is ready. \nIf you want to exit, you will input exit. \nWhen you start, I will recode time. \nAre you ready? (Y/N)")
		for scanner.Scan() {
			line := scanner.Text()

			if line == "N" {
				fmt.Print("Ok! See you.")
				os.Exit(0)
			}

			if line == "exit" {
				os.Exit(0)
			}

			if line == "check"{
				endSecond := time.Now()

				timeDiff := endSecond.Sub(startSecond)

				totalMin := int(timeDiff.Minutes())% 60
				totalSec := int(timeDiff.Seconds())% 60
				fmt.Printf("Time Diff == %d:%d \n",totalMin,totalSec)

			}

			//fmt.Println(line) // Println will add back the final '\n'
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}else{
		log.Fatal("Making number is fail.")
		os.Exit(-1)
	}

	fmt.Println(arrComputerNumbers)
}
