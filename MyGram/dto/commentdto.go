package dto

type Comment struct {
	Message string `validate:"required"`
	PhotoID uint   `validate:"required"`
}

type CommentMessage struct {
	Message string `validate:"required"`
}
