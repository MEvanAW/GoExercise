package database

import (
	"log"
	"time"

	"example.id/mygram/dto"
	"example.id/mygram/models"
)

func CreateSocialMedia(userID uint, socmedDto *dto.SocialMedia) (models.SocialMedia, error) {
	if db == nil {
		return models.SocialMedia{}, ErrDbNotStarted
	}
	newSocmed := models.SocialMedia{
		UserID:         userID,
		Name:           socmedDto.Name,
		SocialMediaUrl: socmedDto.SocialMediaUrl,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	if err := db.Create(&newSocmed).Error; err != nil {
		return models.SocialMedia{}, err
	}
	log.Printf("Social Media Created: %+v\n", newSocmed)
	return newSocmed, nil
}
