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

func GetAllSocialMedias(ctx *gin.Context) {
	socmeds, err := database.GetAllSocialMedias()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	socmedsResponse := make([]responses.GetSocialMedia, len(socmeds))
	userDtos := make(map[uint]dto.UserUpdate)
	for i, socmed := range socmeds {
		socmedsResponse[i].Set(socmed)
		userDto, ok := userDtos[socmed.UserID]
		if !ok {
			userDto, err = database.GetUsernameAndEmail(socmed.UserID)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			userDtos[socmed.UserID] = userDto
		}
		socmedsResponse[i].User = responses.UserSocialMedia{
			ID:       socmed.UserID,
			Username: userDto.Username,
		}
	}
	ctx.JSON(http.StatusOK, responses.GetAllSocialMedias{
		SocialMedias: socmedsResponse,
	})
}
