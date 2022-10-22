package database

import (
	"log"
	"time"

	"example.id/mygram/dto"
	_ "example.id/mygram/dto"
	"example.id/mygram/models"
)

func CreatePhoto(userID uint, photoDto *dto.Photo) error {
	if db == nil {
		return ErrDbNotStarted
	}
	newPhoto := models.Photo{
		Title:    photoDto.Title,
		Caption:  photoDto.Caption,
		PhotoUrl: photoDto.PhotoUrl,
		UserID:   userID,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := db.Create(&newPhoto).Error; err != nil {
		return err
	}
	log.Println("Photo Created:", newPhoto)
	return nil
}
