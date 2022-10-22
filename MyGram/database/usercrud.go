package database

import (
	"errors"
	"log"
	"time"

	"example.id/mygram/dto"
	"example.id/mygram/models"
	"golang.org/x/crypto/bcrypt"
)

var ErrDbNotStarted error = errors.New("DB hasn't started yet.")

func CreateUser(userRegister *dto.UserRegister) error {
	if db == nil {
		return ErrDbNotStarted
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

func UpdateUser(id uint, userDto *dto.UserUpdate) error {
	user, err := getUserWithoutPreload(id)
	if err != nil {
		return err
	}
	if userDto.Email != "" {
		user.Email = userDto.Email
	}
	if userDto.Username != "" {
		user.Username = userDto.Username
	}
	err = db.Save(&user).Error
	if err != nil {
		return err
	}
	log.Printf("User Updated: %+v\n", user)
	return nil
}

func getUserWithoutPreload(id uint) (models.User, error) {
	user := models.User{}
	if db == nil {
		return user, ErrDbNotStarted
	}
	err := db.Model(&models.User{}).Take(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
