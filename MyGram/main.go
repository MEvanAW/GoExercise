package main

import (
	"log"

	"example.id/mygram/database"
	"example.id/mygram/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	database.StartDB()
	// CREATE USER
	bytes, err := bcrypt.GenerateFromPassword([]byte("klaten"), 4)
	if err != nil {
		log.Println(err.Error())
		return
	}
	newUser := models.User{
		Username: "evan60031",
		Email:    "evan60031@gmail.com",
		Password: string(bytes),
		Age:      23,
	}
	err = database.CreateUser(&newUser)
	if err != nil {
		log.Println(err.Error())
	}
}
