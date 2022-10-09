package database

import (
	"fmt"
	"log"

	"excercise.id/orderapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = "localhost"
	user   = "postgres"
	dbPort = "5432"
	dbName = "order-api"
	db     *gorm.DB
	err    error
)

func StartDB() {
	var password string
	fmt.Print("Enter db password: ")
	fmt.Scanln(&password)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}
