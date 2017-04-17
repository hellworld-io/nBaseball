package records

import (
	"testing"
	"path"
	"fmt"
)

func TestReadJsonFile(test *testing.T){
	dataFilePath := path.Join("../data", "UserRecord.json")

	userRecord := new(UserRecord)

	userRecord.readUserRecord(dataFilePath)

	fmt.Println("userRecord =======> ",userRecord.Records[0])
}
