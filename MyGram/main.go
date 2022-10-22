package main

import (
	"log"

	"example.id/mygram/database"
	"example.id/mygram/dto"
)

func main() {
	database.StartDB()
	// CREATE USER
	// userRegister := dto.UserRegister{
	// 	Username: "naufal",
	// 	Email:    "naufal@gmail.com",
	// 	Password: "aduhlali",
	// 	Age:      23,
	// }
	// err := database.CreateUser(&userRegister)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// UPDATE USER
	userUpdate := dto.UserUpdate{
		Username: "m.evan.aw",
		Email:    "m.evan.aw@gmail.com",
	}
	err := database.UpdateUser(2, &userUpdate)
	if err != nil {
		log.Println(err.Error())
	}
	// DELETE USER
	// if err := database.DeleteUserById(3); err != nil {
	// 	log.Println(err.Error())
	// }
}
