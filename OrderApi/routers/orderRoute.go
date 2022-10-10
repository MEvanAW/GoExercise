package routers

import (
	"excercise.id/orderapi/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders/:orderID", controllers.GetOrder)
	return router
}
