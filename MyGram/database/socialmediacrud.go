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

func GetAllSocialMedias() ([]models.SocialMedia, error) {
	socmeds := make([]models.SocialMedia, 1)
	if err := db.Model(&models.SocialMedia{}).Find(&socmeds).Error; err != nil {
		return nil, err
	}
	log.Println("All social medias is read from db.")
	return socmeds, nil
}
