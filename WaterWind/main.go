package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"example.id/waterwind/models"
)

func main() {
	jsonFile, err := os.Open("waterwind.json")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Successfully opened waterwind.json")
	defer jsonFile.Close()
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err.Error())
	}
	var status models.Status
	json.Unmarshal(bytes, &status)
	log.Println(status)
}
