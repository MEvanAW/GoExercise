package database

import (
	"log"
	"time"

	"example.id/mygram/dto"
	"example.id/mygram/models"
)

func CreateComment(userID uint, commentDto *dto.Comment) error {
	if db == nil {
		return ErrDbNotStarted
	}
	newComment := models.Comment{
		UserID:  userID,
		PhotoID: commentDto.PhotoID,
		Message: commentDto.Message,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := db.Create(&newComment).Error; err != nil {
		return err
	}
	log.Printf("Photo Created: %+v\n", newComment)
	return nil
}
