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
	"strconv"
)

var arrComputerNumbers []int
var iNumberMaxLength = flag.Int("maxL", 3, "must create 3 to 5 number length")
var strCreateNumberOption = flag.String("lOpt", "1", "Option \n 1. First number does no use 0 only. \n 2. All number do not use 0. \n 3. 0 to 9 numbers use")
var strResult string

func main() {
	flag.Parse()

	flag.Args()

	//fmt.Println(flag.Args())
	//fmt.Println("*iNumberMaxLength",*iNumberMaxLength)


	if(!util.CheckArgNumberLength(*iNumberMaxLength)){
		log.Fatal("Making number length is 3 to 5.")
		os.Exit(-1)
	}

	arrComputerNumbers = numbers.MakeRandomNumbersByComputer(*iNumberMaxLength, *strCreateNumberOption)

	if (arrComputerNumbers != nil){
		startSecond := time.Now()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Computer is ready. \nIf you want to exit, you will input exit. \nWhen you start, It will be recoded time.\n")
		for scanner.Scan() {
			line := scanner.Text()

			if line == "exit" {
				os.Exit(0)
			}

			iLine, err := strconv.Atoi(line);
			if err != nil {
				fmt.Printf("%q does not looks like a number.\n", line)
			}else{
				if (!util.CheckUserNumberLength(line, *iNumberMaxLength)){
					fmt.Printf("Max number length is %d\n",*iNumberMaxLength)
				}else{
					strResult = numbers.CheckUserNumber(iLine, arrComputerNumbers)

					if(strResult == "Out"){

					}else{
						fmt.Printf("%q result is %q\n", iLine, strResult)
					}
				}
			}

			if line == "check"{
				endSecond := time.Now()

				timeDiff := endSecond.Sub(startSecond)

				totalMin := int(timeDiff.Minutes())% 60
				totalSec := int(timeDiff.Seconds())% 60
				fmt.Printf("Time Diff == %d:%d \n",totalMin,totalSec)

			}
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
