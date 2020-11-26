package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetJSONData(data interface{}) {
	jsonFile, err := os.Open("database/布袋戲資料.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &data)
}
