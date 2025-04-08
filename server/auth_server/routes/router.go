package routes

import (
	"auth-service/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine){
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}