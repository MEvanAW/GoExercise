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

func GetAllPhotos() ([]models.Photo, error) {
	photos := make([]models.Photo, 1)
	if err := db.Model(&models.Photo{}).Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

func GetSinglePhoto(photoID uint) (models.Photo, error) {
	photo := models.Photo{}
	if db == nil {
		return photo, ErrDbNotStarted
	}
	err := db.Model(&models.Photo{}).Take(&photo, photoID).Error
	return photo, err
}

func UpdatePhoto(photoID uint, photoDto *dto.Photo) (UpdatedAt time.Time, err error) {
	photo, err := GetSinglePhoto(photoID)
	if err != nil {
		var temp time.Time
		return temp, err
	}
	if photoDto.Title != "" {
		photo.Title = photoDto.Title
	}
	if photoDto.PhotoUrl != "" {
		photo.PhotoUrl = photoDto.PhotoUrl
	}
	photo.Caption = photoDto.Caption
	photo.UpdatedAt = time.Now()
	err = db.Save(&photo).Error
	if err == nil {
		log.Printf("Photo Updated: %+v\n", photo)
	} else {
		log.Printf("Failed Update Photo %+v with %+v because of %q\n", photo, photoDto, err.Error())
	}
	return photo.UpdatedAt, err
}

func DeletePhotoById(photoID uint) error {
	photo, err := GetSinglePhoto(photoID)
	if err != nil {
		return err
	}
	if err := db.Delete(&photo, photoID).Error; err != nil {
		return err
	}
	log.Println("Photo with ID", photoID, "has been successfully deleted.")
	return nil
}
