package controllers

import (
	"errors"
	"net/http"

	"example.id/mygram/database"
	"example.id/mygram/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
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
		validationAbort(err, ctx)
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
	ctx.JSON(http.StatusCreated, newUser)
}

func LoginUser(ctx *gin.Context) {
	var userLogin dto.UserLogin
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			errorMessageStr: err.Error(),
		})
		return
	}
	if err := validate.Struct(&userLogin); err != nil {
		validationAbort(err, ctx)
		return
	}
	jwt, err := database.GenerateToken(userLogin)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, database.ErrPasswordMismatch) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				errorMessageStr: "Email or password is incorrect.",
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}

func validationAbort(err error, ctx *gin.Context) {
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
}
