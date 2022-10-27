package controllers

import (
	"errors"
	"net/http"

	"example.id/mygram/database"
	"example.id/mygram/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgconn"
)

var validate = validator.New()

const (
	messageStr         = "message"
	errorMessageStr    = "error_message"
	uniqueViolationErr = "23505"
)

func RegisterUser(ctx *gin.Context) {
	var newUser dto.UserRegister
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			errorMessageStr: err.Error(),
		})
		return
	}
	if err := validate.Struct(&newUser); err != nil {
		var errorMessage string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage += err.Error() + "\n"
		}
		if len(errorMessage) > 0 {
			errorMessage = errorMessage[:len(errorMessage)-1]
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			errorMessageStr: err.Error(),
		})
		return
	}
	err := database.CreateUser(&newUser)
	if err != nil {
		var perr *pgconn.PgError
		if ok := errors.As(err, &perr); ok {
			if perr.Code == uniqueViolationErr {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					messageStr: "The email or username is already registered. If it is yours, do login instead.",
				})
				return
			}
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user": newUser,
	})
}
