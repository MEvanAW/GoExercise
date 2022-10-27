package responses

import (
	"time"

	"example.id/mygram/models"
)

type CreateSocialMedia struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type GetAllSocialMedias struct {
	SocialMedias []GetSocialMedia `json:"social_medias"`
}

type GetSocialMedia struct {
	models.SocialMedia
	User UserSocialMedia
}

type UserSocialMedia struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func (getSocmed *GetSocialMedia) Set(socmed models.SocialMedia) {
	getSocmed.ID = socmed.ID
	getSocmed.CreatedAt = socmed.CreatedAt
	getSocmed.UpdatedAt = socmed.UpdatedAt
	getSocmed.Name = socmed.Name
	getSocmed.SocialMediaUrl = socmed.SocialMediaUrl
	getSocmed.UserID = socmed.UserID
}
