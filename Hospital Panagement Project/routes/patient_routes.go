package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/hospital-management-portal/controllers"
	"github.com/yourusername/hospital-management-portal/middleware"
)

func PatientRoutes(router *gin.Engine) {
	patientRoutes := router.Group("/patient")
	patientRoutes.Use(middleware.AuthMiddleware())
	patientRoutes.Use(middleware.RoleMiddleware("Patient"))
	{
		patientRoutes.POST("/appointments", controllers.ScheduleAppointment)
		patientRoutes.GET("/doctors", controllers.ViewAvailableDoctors)
		patientRoutes.GET("/appointments", controllers.ViewAppointments)
	}
}
