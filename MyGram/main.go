package main

import (
	"log"

	"example.id/mygram/database"
	_ "example.id/mygram/dto"
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
	// userUpdate := dto.UserUpdate{
	// 	Username: "evan60031",
	// 	Email:    "evan60031@gmail.com",
	// }
	// err := database.UpdateUser(2, &userUpdate)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// DELETE USER
	// if err := database.DeleteUserById(3); err != nil {
	// 	log.Println(err.Error())
	// }
	// CREATE PHOTO
	// photoDto := dto.Photo{
	// 	Title:    "Linkedin Photo Profile",
	// 	Caption:  "UI/UX Designer at PT Javan Cipta Solusi.",
	// 	PhotoUrl: "https://media-exp1.licdn.com/dms/image/C5603AQH5R3vCtuyvUg/profile-displayphoto-shrink_400_400/0/1622972541475?e=1672272000&v=beta&t=139NJp5PNckox4pix6JynXuAg9QqGgomYkYNeNLGOtY",
	// }
	// if err := database.CreatePhoto(1, &photoDto); err != nil {
	// 	log.Println(err.Error())
	// }
	// GET ALL PHOTO
	photos, err := database.GetAllPhotos()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("%+v\n", photos)
	}
}
