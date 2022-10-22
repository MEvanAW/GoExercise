package main

import (
	"log"

	"example.id/mygram/database"
	"example.id/mygram/dto"
)

func main() {
	database.StartDB()
	// CREATE USER
	userRegister := dto.UserRegister{
		Username: "fandi",
		Email:    "fandi@gmail.com",
		Password: "qwerty",
		Age:      23,
	}
	err := database.CreateUser(&userRegister)
	if err != nil {
		log.Println(err.Error())
	}
}
