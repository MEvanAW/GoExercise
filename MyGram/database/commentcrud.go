package database

import (
	"log"
	"time"

	"example.id/mygram/dto"
	"example.id/mygram/models"
)

func CreateComment(userID uint, commentDto *dto.Comment) (models.Comment, error) {
	if db == nil {
		return models.Comment{}, ErrDbNotStarted
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
		return models.Comment{}, err
	}
	log.Printf("Photo Created: %+v\n", newComment)
	return newComment, nil
}

func GetAllComments() ([]models.Comment, error) {
	comments := make([]models.Comment, 1)
	if err := db.Model(&models.Comment{}).Find(&comments).Error; err != nil {
		return nil, err
	}
	log.Println("All comments is read from db.")
	return comments, nil
}
