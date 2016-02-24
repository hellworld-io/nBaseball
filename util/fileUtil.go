package util

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
)

var JsonData []map[string]interface{}

type UserResult struct {
	UserResults []struct{
		No string `json:"no"`
		Name string `json:"name"`
		Time string `json:"time"`
		Try string `json:"try"`
	} `json:"result"`
}


func (u *UserResult) ReadJsonFile(strFileName string) {
	file, err := ioutil.ReadFile(strFileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	userResultJsonData := &UserResult{}

	err = json.Unmarshal(file, &userResultJsonData)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
