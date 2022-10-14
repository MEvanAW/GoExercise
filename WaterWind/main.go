package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"example.id/waterwind/models"
	"github.com/gin-gonic/gin"
)

func main() {
	go jsonUpdate()
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", func(c *gin.Context) {
		jsonFile, err := os.Open("waterwind.json")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			log.Println(err.Error())
			return
		}
		bytes, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			log.Println(err.Error())
			return
		}
		var status models.Status
		json.Unmarshal(bytes, &status)
		log.Printf("in router: %+v\n", status)
		waterStatus := ""
		windStatus := ""
		switch {
		case status.Status.Water <= 5:
			waterStatus = "safe"
		case status.Status.Water <= 8:
			waterStatus = "alert"
		default:
			waterStatus = "danger"
		}
		switch {
		case status.Status.Wind <= 6:
			windStatus = "safe"
		case status.Status.Wind <= 15:
			windStatus = "alert"
		default:
			windStatus = "danger"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"water":       status.Status.Water,
			"waterStatus": waterStatus,
			"wind":        status.Status.Wind,
			"windStatus":  windStatus,
		})
	})
	router.Run("localhost:8080")
}

func jsonUpdate() {
	for {
		waterStatus, windStatus := rand.Intn(3), rand.Intn(3)
		var water, wind int
		switch waterStatus {
		case 0:
			water = rand.Intn(5) + 1
		case 1:
			water = rand.Intn(3) + 6
		default:
			water = rand.Intn(92) + 9
		}
		switch windStatus {
		case 0:
			wind = rand.Intn(6) + 1
		case 1:
			wind = rand.Intn(9) + 7
		default:
			wind = rand.Intn(85) + 16
		}
		jsonString := fmt.Sprintf(`{
	"status": {
		"water": %d,
		"wind": %d
	}
}`, water, wind)
		jsonFile, err := os.Create("waterwind.json")
		if err != nil {
			log.Println("Error in jsonUpdate:", err.Error())
			continue
		}
		_, err = jsonFile.Write([]byte(jsonString))
		if err != nil {
			log.Println("Error in jsonUpdate:", err.Error())
			continue
		}
		log.Printf("in jsonUpdater: {Status:{Water:%d Wind:%d}}\n", water, wind)
		time.Sleep(15 * time.Second)
	}
}
