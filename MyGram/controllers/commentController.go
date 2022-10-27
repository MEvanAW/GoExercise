package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"example.id/mygram/controllers/responses"
	"example.id/mygram/database"
	"example.id/mygram/dto"
	"example.id/mygram/models"
	"example.id/mygram/utils/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(ctx *gin.Context) {
	var newComment dto.Comment
	if err := ctx.ShouldBindJSON(&newComment); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&newComment); err != nil {
		validationAbort(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	comment, err := database.CreateComment(userID, &newComment)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, responses.CreateComment{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    userID,
		CreatedAt: comment.CreatedAt,
	})
}

func GetAllComments(ctx *gin.Context) {
	comments, err := database.GetAllComments()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	commentsResponse := make([]responses.GetComment, len(comments))
	users := make(map[uint]models.User)
	photos := make(map[uint]models.Photo)
	for i, comment := range comments {
		commentsResponse[i].Set(comment)
		user, ok := users[comment.UserID]
		if !ok {
			user, err = database.GetUserWithoutPreload(comment.UserID)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			users[comment.UserID] = user
		}
		commentsResponse[i].User = responses.UserComment{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		}
		photo, ok := photos[comment.PhotoID]
		if !ok {
			photo, err = database.GetSinglePhoto(comment.PhotoID)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			photos[comment.PhotoID] = photo
		}
		commentsResponse[i].Photo = photo
	}
	ctx.JSON(http.StatusOK, commentsResponse)
}

func UpdateComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")
	parsedID, err := strconv.ParseUint(commentID, 10, 0)
	if err != nil {
		abortBadRequest(err, ctx)
		return
	}
	var commentDto dto.CommentMessage
	if err := ctx.ShouldBindJSON(&commentDto); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&commentDto); err != nil {
		validationAbort(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	comment, err := database.UpdateComment(uint(parsedID), userID, &commentDto)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, responses.ErrorMessage{
				ErrorMessage: fmt.Sprintf("Comment with ID %d is not found.", parsedID),
			})
			return
		}
		if errors.Is(err, database.ErrIllegalUpdate) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, responses.ErrorMessage{
				ErrorMessage: err.Error(),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, comment)
}
