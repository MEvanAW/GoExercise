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

var ErrNotFound error = errors.New("Order with the ID provided is not found.")

// CreateOrder godoc
// @Summary      Create an order
// @Description  Create an order including its items, if provided.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        order body models.Order true "JSON of the order to be made. Please remove both 'id' and 'orderID' line."
// @Success      201  {object}  models.Order
// @Failure      400  {object}  string
// @Failure      500  {object}  nil
// @Router       /orders [post]
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

// GetOrder godoc
// @Summary      Get an order
// @Description  get order by ID
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "orderID"
// @Success      200  {object}  models.Order
// @Failure      400  {object}  nil
// @Failure      404  {object}  string
// @Failure      500  {object}  nil
// @Router       /orders/{orderID} [get]
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
				"error_message": fmt.Sprintf("Order with ID %d is not found.", parsedID),
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

// UpdateOrder godoc
// @Summary      Update an order
// @Description  update order by ID including its items. Previous items are discarded.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "orderID"
// @Param        order body models.Order true "JSON of the order to be made. Please remove the 'orderID' line."
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  nil
// @Router       /orders/{orderID} [put]
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_message": fmt.Sprintf("Order with ID %d is not found.", parsedID),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with id %d has been successfully updated.", parsedID),
	})
}

// DeleteOrder godoc
// @Summary      Delete an order
// @Description  delete order by ID including its items.
// @Tags         orders
// @Accept       json
// @Produce      json
// @Param        id   path      uint  true  "orderID"
// @Success      200  {object}  string
// @Failure      400  {object}  nil
// @Failure      404  {object}  string
// @Failure      500  {object}  nil
// @Router       /orders/{orderID} [delete]
func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	parsedID, err := strconv.ParseUint(orderID, 10, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := database.DeleteOrderById(uint(parsedID)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_message": fmt.Sprintf("Order with ID %d is not found.", parsedID),
			})
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Order with id %d has been successfully deleted.", parsedID),
	})
}
