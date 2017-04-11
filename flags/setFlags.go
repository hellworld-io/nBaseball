package flags

import (
	"flag"
	"fmt"
	"os"
)

const usages string = `Usage of nBaseball:
    start -max int
        create 3 to 5 number of length ex) nBaseball start -max 3
    start -option string
        choose option
        1. All number do not use 0.
        2. First number does no use 0 only.
        3. 0 to 9 numbers use
	`

var (
	startCmd   = flag.NewFlagSet("start", flag.ExitOnError)
	maxNumberLength  = startCmd.Int("max", 0, "create 3 to 5 number of length ex) nBaseball start -max 3")
	numberOption = startCmd.String("option", "", "Option \n 1. All number do not use 0. \n 2. First number does no use 0 only. \n 3. 0 to 9 numbers use")
)

func defaultMessage() {
	fmt.Println(usages)
	flag.PrintDefaults()
}

func checkFlags(maxLen int, option string) bool{
	if maxLen == 0 {
		return false
	}

	if option == "" {
		return false
	}

	if maxLen < 3 || maxLen > 5 {
		return false
	}

	if option != "1" && option != "2" && option != "3" {
		return false
	}

	return true
}

func SetFlags() (maxArg int, optionArg string) {

	flag.Usage = defaultMessage

	if len(os.Args) < 2 {
		defaultMessage()
		os.Exit(-1)
	}

	if os.Args[1] == "start" {
		startCmd.Parse(os.Args[2:])
	} else {
		defaultMessage()
		os.Exit(-1)
	}

	if startCmd.Parsed() {
		if !checkFlags(*maxNumberLength, *numberOption) {
			startCmd.PrintDefaults()
			os.Exit(-1)
		}

		maxArg = *maxNumberLength
		optionArg = *numberOption
	}

	return maxArg, optionArg
}