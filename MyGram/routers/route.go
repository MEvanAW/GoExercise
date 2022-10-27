package routers

import (
	"example.id/mygram/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("users/register", controllers.RegisterUser)
	return router
}
