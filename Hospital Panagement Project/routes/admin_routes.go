package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/hospital-management-portal/controllers"
	"github.com/yourusername/hospital-management-portal/middleware"
)

func AdminRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware())
	adminRoutes.Use(middleware.RoleMiddleware("Admin"))
	{
		adminRoutes.POST("/doctors", controllers.CreateDoctor)
		adminRoutes.PUT("/doctors/:doctorId", controllers.UpdateDoctor)
		adminRoutes.DELETE("/doctors/:doctorId", controllers.DeleteDoctor)
		adminRoutes.GET("/doctors", controllers.ListDoctors)
		adminRoutes.POST("/patients", controllers.CreatePatient)
		adminRoutes.PUT("/patients/:patientId", controllers.UpdatePatient)
		adminRoutes.DELETE("/patients/:patientId", controllers.DeletePatient)
		adminRoutes.GET("/patients", controllers.ListPatients)
		adminRoutes.GET("/appointments", controllers.ViewAppointments)
	}
}
