package database

import (
	"errors"
	"log"
	"time"

	"example.id/mygram/dto"
	"example.id/mygram/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(userRegister *dto.UserRegister) error {
	if db == nil {
		return errors.New("DB hasn't started yet.")
	}
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 4)
	if err != nil {
		return err
	}
	newUser := models.User{
		Username: userRegister.Username,
		Email:    userRegister.Email,
		Password: string(passwordBytes),
		Age:      userRegister.Age,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	err = db.Create(&newUser).Error
	if err != nil {
		return err
	}
	log.Println("User Created:", newUser)
	return nil
}
