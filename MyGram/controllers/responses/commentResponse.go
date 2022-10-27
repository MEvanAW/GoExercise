package responses

import (
	"time"

	"example.id/mygram/models"
)

type CreateComment struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetComment struct {
	CreateComment
	UpdatedAt time.Time `json:"updated_at"`
	User      UserComment
	Photo     models.Photo
}

type UserComment struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (getComment *GetComment) Set(comment models.Comment) {
	getComment.ID = comment.ID
	getComment.CreatedAt = comment.CreatedAt
	getComment.UpdatedAt = comment.UpdatedAt
	getComment.UserID = comment.UserID
	getComment.PhotoID = comment.PhotoID
	getComment.Message = comment.Message
}
