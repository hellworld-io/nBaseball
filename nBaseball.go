/**
    file name : nBaseball.go
 */
package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"nBaseball/util"
	"nBaseball/numbers"
	"time"
	"strconv"
	"bufio"
)

var arrComputerNumbers = []int{1,2,3}	//Test data
var iNumberMaxLength = flag.Int("maxL", 3, "must create 3 to 5 number length")
var strCreateNumberOption = flag.String("lOpt", "1", "Option \n 1. First number does no use 0 only. \n 2. All number do not use 0. \n 3. 0 to 9 numbers use")
var strResult string
var bClear = bool(true)

func main() {
	flag.Parse()

	flag.Args()

	//fmt.Println(flag.Args())
	//fmt.Println("*iNumberMaxLength",*iNumberMaxLength)


	if(!util.CheckArgNumberLength(*iNumberMaxLength)){
		log.Fatal("Making number length is 3 to 5.")
		os.Exit(-1)
	}

	//[TODO] for test temporary comment
	//numbers.MakeRandomNumbersByComputer(*iNumberMaxLength, *strCreateNumberOption)
	//numbers.ArrRandomNumbers
	if (arrComputerNumbers != nil){
		startSecond := time.Now()

		fmt.Print("Computer is ready. \nIf you want to exit, you will input exit. \nWhen you start, It will be recoded time.\n")

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Computer is ready. \nIf you want to exit, you will input exit. \nWhen you start, It will be recoded time.\n")

		for scanner.Scan() {

			line := scanner.Text()


			if line == "exit" {
				os.Exit(0)
			}

			if(bClear){
				if _, err := strconv.Atoi(line); err != nil {
					fmt.Printf("%q does not looks like a number.\n", line)
				}else{
					if (!util.CheckUserNumberLength(line, *iNumberMaxLength)){
						fmt.Printf("Max number length is %d\n",*iNumberMaxLength)
					}else{
						if(util.CheckUserInputNumber(line)){
							strResult = numbers.CompareUserNumbersAndRandomNumbers(line, arrComputerNumbers)

							if(strResult == "Out"){
								endSecond := time.Now()

								timeDiff := endSecond.Sub(startSecond)

								hours0 := int(timeDiff.Hours())
								days := hours0 / 24
								hours := hours0 % 24
								totalMin := int(timeDiff.Minutes())% 60
								totalSec := int(timeDiff.Seconds())% 60
								fmt.Printf("Your result is %d days + %d hours + %d minutes + %d seconds \n", days, hours, totalMin, totalSec)
								fmt.Println("If you want to see results, press enter.")
								//os.Exit(-1)

								bClear = false
							}else{
								fmt.Printf("%q result is %q\n", line, strResult)
							}
						}else{
							fmt.Println("Do not use duplication number")
						}
					}
				}
			}else{
				fmt.Println("Here is currently results.")
				var strFileName = string( "./data/UserResult.json")
				util.ReadJsonFile(strFileName)
				fmt.Println(util.JsonData)
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
