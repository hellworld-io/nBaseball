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
	_"encoding/json"
)

var arrComputerNumbers = []int{1,2,3}	//Test data
var iNumberMaxLength = flag.Int("maxL", 3, "must create 3 to 5 number length")
var strCreateNumberOption = flag.String("lOpt", "1", "Option \n 1. All number do not use 0. \n 2. First number does no use 0 only. \n 3. 0 to 9 numbers use")
var strResult string
var bClear = bool(false)
var iTryTotalCount int


func showResults(tStartTime time.Time, tEndTime time.Time){
	timeDiff := tEndTime.Sub(tStartTime)

	hours0 := int(timeDiff.Hours())
	days := hours0 / 24
	hours := hours0 % 24
	totalMin := int(timeDiff.Minutes())% 60
	totalSec := int(timeDiff.Seconds())% 60
	fmt.Printf("Your result is %dD %dH %dM %dS, and you try %d\n", days, hours, totalMin, totalSec, iTryTotalCount)

	//fmt.Println("If you want to see results, press enter.")
}

func setResult() (result []string){
	result = make([]string, 5)
	for i := range result {
		result[i] = "│"
	}

	result = append([]string{
		"                       ┌────────────┐                           ",
		"                       │   Result   │                           ",
		"                       └────────────┘                           ",
		"┌────┬────────────────────────┬──────────────────┬─────────────┐",
		"│ No │          Name          │       Time       │     Try     │",
		"├────┼────────────────────────┼──────────────────┼─────────────┤"},
		result...)
	return append(result,
		"└────┴────────────────────────┴──────────────────┴─────────────┘")
}

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
	//arrComputerNumbers
	//numbers.ArrRandomNumbers

	if (arrComputerNumbers != nil){
		startSecond := time.Now()

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Computer is ready. \nIf you want to exit, you will input exit. \nWhen you started, It will be recoded time.\n")

		for scanner.Scan() {
			line := scanner.Text()

			if line == "exit" {
				os.Exit(0)
			}

			if(!bClear){
				iTryTotalCount++
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

								showResults(startSecond, endSecond)
								//os.Exit(-1)

								bClear = true
							}else{
								fmt.Printf("%q result is %q\n", line, strResult)
							}
						}else{
							fmt.Println("Do not use duplication number")
						}
					}
				}
			}else{
				//fmt.Println(setResult())

				for _, val := range setResult() {
					fmt.Fprintln(os.Stdout, val)
				}

				//var strFileName = string( "./data/UserResult.json")
				//util.ReadJsonFile(strFileName)
				//var uinfo = &util.UserResult{}
				//fmt.Println(util.UserResult{})
				//fmt.Println(uinfo)
				/*
				fmt.Println("Here is currently results.")
				var strFileName = string( "./data/UserResult.json")
				util.ReadJsonFile(strFileName)
				fmt.Println(util.JsonData[0]["Name"].(string))
				util.JsonData[0]["Name"] = "aaa"
				fmt.Println(util.JsonData[0]["Name"].(string))
				fmt.Println(util.JsonData)
				jsonUserResult, err := json.Marshal(util.JsonData)

				if err != nil {
					log.Fatal("Making json file error!!! ")
					os.Exit(-1)
				}


				jsonUpdate, err := os.Create("./data/UserResult.json")

				if err != nil {
					log.Fatal("Making json file error!!! ")
					os.Exit(-1)
				}
				//defer jsonUpdate.Close()

				jsonUpdate.Write(jsonUserResult)
				jsonUpdate.Close()
				*/
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}else{
		log.Fatal("Making number is fail.")
		os.Exit(-1)
	}

	//fmt.Println(arrComputerNumbers)
}
