package controllers

import (
	"net/http"
	"time"

	"example.id/mygram/controllers/responses"
	"example.id/mygram/database"
	"example.id/mygram/dto"
	"example.id/mygram/utils/token"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(ctx *gin.Context) {
	var newPhoto dto.Photo
	if err := ctx.ShouldBindJSON(&newPhoto); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&newPhoto); err != nil {
		validationAbort(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ID, err := database.CreatePhoto(userID, &newPhoto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, responses.CreatePhoto{
		ID:        ID,
		Title:     newPhoto.Title,
		Caption:   newPhoto.Caption,
		PhotoUrl:  newPhoto.PhotoUrl,
		UserID:    userID,
		CreatedAt: time.Now(),
	})
}
