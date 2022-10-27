package controllers

import (
	"net/http"

	"example.id/mygram/controllers/responses"
	"example.id/mygram/database"
	"example.id/mygram/dto"
	"example.id/mygram/utils/token"
	"github.com/gin-gonic/gin"
)

func CreateSocialMedia(ctx *gin.Context) {
	var newSocmed dto.SocialMedia
	if err := ctx.ShouldBindJSON(&newSocmed); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&newSocmed); err != nil {
		validationAbort(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	socmed, err := database.CreateSocialMedia(userID, &newSocmed)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, responses.CreateSocialMedia{
		ID:             socmed.ID,
		Name:           socmed.Name,
		SocialMediaUrl: socmed.SocialMediaUrl,
		UserID:         userID,
		CreatedAt:      socmed.CreatedAt,
	})
}
