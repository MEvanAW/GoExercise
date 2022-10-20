package database

import (
	"errors"
	"log"
	"time"

	"example.id/mygram/models"
)

func CreateUser(user *models.User) error {
	if db == nil {
		return errors.New("DB hasn't started yet.")
	}
	var zero time.Time
	if user.CreatedAt == zero {
		user.CreatedAt = time.Now()
	}
	if user.UpdatedAt == zero {
		user.UpdatedAt = time.Now()
	}
	err = db.Create(user).Error
	if err != nil {
		return err
	}
	log.Println("User Created:", user)
	return nil
}
