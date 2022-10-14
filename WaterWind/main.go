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
		c.HTML(http.StatusOK, "index.html", gin.H{
			"water": status.Status.Water,
			"wind":  status.Status.Wind,
		})
	})
	router.Run("localhost:8080")
}

func jsonUpdate() {
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
