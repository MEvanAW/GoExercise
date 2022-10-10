package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"excercise.id/orderapi/database"
	"excercise.id/orderapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(ctx *gin.Context) {
	var newOrder models.Order
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := database.CreateOrder(&newOrder)
	if err != nil {
		if errors.Is(err, models.ErrItemCodeEmpty) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			return
		}
		if errors.Is(err, models.ErrCustomerNameEmpty) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error_message": err.Error(),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

func GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	var orderData models.Order
	parsedID, err := strconv.ParseUint(orderID, 10, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	orderData, err = database.GetOrderById(uint(parsedID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_message": fmt.Sprintf("Order with ID %d not found.", parsedID),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"order": orderData,
	})
}

func UpdateOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	parsedID, err := strconv.ParseUint(orderID, 10, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var updatedOrder models.Order
	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := database.UpdateOrderById(uint(parsedID), &updatedOrder); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with id %d has been successfully updated.", parsedID),
	})
}
