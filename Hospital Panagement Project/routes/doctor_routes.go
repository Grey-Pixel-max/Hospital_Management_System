package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/hospital-management-portal/controllers"
	"github.com/yourusername/hospital-management-portal/middleware"
)

func DoctorRoutes(router *gin.Engine) {
	doctorRoutes := router.Group("/doctor")
	doctorRoutes.Use(middleware.AuthMiddleware())
	doctorRoutes.Use(middleware.RoleMiddleware("Doctor"))
	{
		doctorRoutes.PUT("/patients/:patientId", controllers.UpdatePatient)
		doctorRoutes.DELETE("/patients/:patientId", controllers.DeletePatient)
		doctorRoutes.GET("/patients", controllers.ListPatients)
		doctorRoutes.GET("/appointments", controllers.ViewAppointments)
		doctorRoutes.POST("/availability", controllers.SetAvailability)
	}
}
