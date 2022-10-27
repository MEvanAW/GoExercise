package routers

import (
	"example.id/mygram/controllers"
	"example.id/mygram/middlewares"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("users/register", controllers.RegisterUser)
	router.POST("users/login", controllers.LoginUser)
	router.PUT("users", middlewares.JwtAuthMiddleware(), controllers.UpdateUser)
	router.DELETE("users", middlewares.JwtAuthMiddleware(), controllers.DeleteUser)
	return router
}
