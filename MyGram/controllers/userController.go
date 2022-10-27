package controllers

import (
	"errors"
	"net/http"

	"example.id/mygram/database"
	"example.id/mygram/dto"
	"example.id/mygram/response"
	"example.id/mygram/utils/token"
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
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&newUser); err != nil {
		validationAbort(err, ctx)
		return
	}
	ID, err := database.CreateUser(&newUser)
	if err != nil {
		var perr *pgconn.PgError
		if ok := errors.As(err, &perr); ok {
			if perr.Code == uniqueViolationErr {
				ctx.AbortWithStatusJSON(http.StatusOK, response.Message{
					Message: "The email or username is already registered. If it is yours, do login instead.",
				})
				return
			}
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, response.UserRegister{
		Age:      newUser.Age,
		Email:    newUser.Email,
		ID:       ID,
		Username: newUser.Username,
	})
}

func LoginUser(ctx *gin.Context) {
	var userLogin dto.UserLogin
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&userLogin); err != nil {
		validationAbort(err, ctx)
		return
	}
	jwt, err := database.GenerateToken(userLogin)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, database.ErrPasswordMismatch) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
				ErrorMessage: "Email or password is incorrect.",
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response.UserLogin{
		Token: jwt,
	})
}

func UpdateUser(ctx *gin.Context) {
	var userDto dto.UserUpdate
	if err := ctx.ShouldBindJSON(&userDto); err != nil {
		abortBadRequest(err, ctx)
		return
	}
	if err := validate.Struct(&userDto); err != nil {
		validationAbort(err, ctx)
		return
	}
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	user, err := database.UpdateUser(userID, &userDto)
	if err != nil {
		var perr *pgconn.PgError
		if ok := errors.As(err, &perr); ok {
			if perr.Code == uniqueViolationErr {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					messageStr: "The email or username is already registered.",
				})
				return
			}
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response.UserUpdate{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	})
}

func DeleteUser(ctx *gin.Context) {
	userID, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := database.DeleteUserById(userID); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response.Message{
		Message: "Your account has been successfully deleted",
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
	abortBadRequest(err, ctx)
}

func abortBadRequest(err error, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response.ErrorMessage{
		ErrorMessage: err.Error(),
	})
}
