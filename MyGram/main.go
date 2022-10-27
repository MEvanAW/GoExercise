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
	// userUpdate := dto.UserUpdate{
	// 	Username: "evan60031",
	// 	Email:    "evan60031@gmail.com",
	// }
	// updatedUser, err := database.UpdateUser(2, &userUpdate)
	// if err != nil {
	// 	log.Println(err.Error())
	// } else {
	//  log.Printf("%+v\n", updatedUser)
	// }
	// DELETE USER
	// if err := database.DeleteUserById(3); err != nil {
	// 	log.Println(err.Error())
	// }
	// CREATE PHOTO
	// photoDto := dto.Photo{
	// 	Title:    "Will be updated.",
	// 	Caption:  "Will be updated.",
	// 	PhotoUrl: "https://will.be.com/updated/image/C5/profile-displayphoto-shrink_400/0/5?e=1&v=beta&t=1",
	// }
	// if err := database.CreatePhoto(1, &photoDto); err != nil {
	// 	log.Println(err.Error())
	// }
	// GET ALL PHOTO
	// photos, err := database.GetAllPhotos()
	// if err != nil {
	// 	log.Println(err.Error())
	// } else {
	// 	log.Printf("%+v\n", photos)
	// }
	// UPDATE A PHOTO
	// photoDto := dto.Photo{
	// 	Title:    "Updated!",
	// 	Caption:  "Updated!",
	// 	PhotoUrl: "https://updated.be.com/image/C?e=1",
	// }
	// updatedAt, err := database.UpdatePhoto(7, &photoDto)
	// if err != nil {
	// 	log.Println(err.Error())
	// } else {
	// 	log.Printf("%+v at %v", photoDto, updatedAt)
	// }
	// DELETE A PHOTO
	// err := database.DeletePhotoById(7)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// CREATE A COMMENT
	// commentDto := dto.Comment{
	// 	Message: "To be deleted.",
	// 	PhotoID: 3,
	// }
	// _, err := database.CreateComment(2, &commentDto)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// GET ALL COMMENTS
	// comments, err := database.GetAllComments()
	// if err != nil {
	// 	log.Println(err.Error())
	// } else {
	// 	log.Printf("All comments: %+v\n", comments)
	// }
	// UPDATE A COMMENT
	// messageDto := dto.CommentMessage{
	// 	Message: ":mbkorangtua: :8ball:",
	// }
	// updatedAt, err := database.UpdateComment(4, &messageDto)
	// if err != nil {
	// 	log.Println(err.Error())
	// } else {
	// 	log.Printf("%+v at %v", messageDto, updatedAt)
	// }
	// DELETE A COMMENT
	// err := database.DeleteCommentById(4)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// CREATE SOCIAL MEDIA
	socmedDto := dto.SocialMedia{
		Name:           "Linkedin",
		SocialMediaUrl: "https://www.linkedin.com/in/muhammad-evan-anindya-wahyuaji-7796aa192/",
	}
	_, err := database.CreateSocialMedia(2, &socmedDto)
	if err != nil {
		log.Println(err.Error())
	}
}
