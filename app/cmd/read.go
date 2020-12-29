package cmd

import (
	"encoding/json"
	"io/ioutil"
	"linked-list/app/models"
	"os"
)

func readData(nv chan int) {
	jsonFile, err := os.Open("./app/data/data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	var nodeValues models.Datas
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &nodeValues)

	for _, val := range nodeValues.Datas {
		nv <- val.Data
	}
	close(nv)
}
