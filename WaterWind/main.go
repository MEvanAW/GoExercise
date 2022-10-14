package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"example.id/waterwind/models"
)

func main() {
	updatedChan := make(chan bool)
	go jsonUpdate(updatedChan)
	for {
		updated := <-updatedChan
		if !updated {
			return
		}
		jsonFile, err := os.Open("waterwind.json")
		if err != nil {
			log.Println(err.Error())
			return
		}
		bytes, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Println(err.Error())
			return
		}
		var status models.Status
		json.Unmarshal(bytes, &status)
		log.Println(status)
	}
}

func jsonUpdate(updated chan bool) {
	for {
		water, wind := rand.Intn(100)+1, rand.Intn(100)+1
		jsonString := fmt.Sprintf(`{
			"status": {
				"water": %d,
				"wind": %d
			}
		}`, water, wind)
		jsonFile, err := os.Create("waterwind.json")
		if err != nil {
			log.Println(err.Error())
			updated <- false
			return
		}
		_, err = jsonFile.Write([]byte(jsonString))
		if err != nil {
			log.Println(err.Error())
			updated <- false
			return
		}
		updated <- true
		time.Sleep(15 * time.Second)
	}
}
