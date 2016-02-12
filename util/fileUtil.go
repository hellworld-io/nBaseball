package util

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
)

var JsonData map[string][]map[string]interface{}

func ReadJsonFile(strFileName string) {
	file, err := ioutil.ReadFile(strFileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}//

	err = json.Unmarshal(file, &JsonData)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
