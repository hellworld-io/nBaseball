package records

import (
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)

type UserRecord struct {
	Records []struct{
		No string `json:"no"`
		Name string `json:"name"`
		Time string `json:"time"`
		Try string `json:"try"`
	} `json:"record"`
}

func (userRecord *UserRecord) readUserRecord(strFileName string) {
	file, err := ioutil.ReadFile(strFileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	err = json.Unmarshal(file, &userRecord)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}