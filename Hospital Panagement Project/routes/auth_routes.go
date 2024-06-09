package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/hospital-management-portal/controllers"
)

func AuthRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}
}
