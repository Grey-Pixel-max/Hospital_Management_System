package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/hospital-management-portal/database"
	"github.com/yourusername/hospital-management-portal/routes"
)

func main() {
	router := gin.Default()

	// Connect to the MongoDB database
	database.ConnectDatabase()

	// Set up the routes
	routes.AuthRoutes(router)
	routes.AdminRoutes(router)
	routes.DoctorRoutes(router)
	routes.PatientRoutes(router)

	// Run the server
	router.Run(":8080")
}
