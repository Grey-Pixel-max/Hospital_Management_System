package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/yourusername/hospital-management-portal/database"
	"github.com/yourusername/hospital-management-portal/models"
)

func ScheduleAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.BindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment.PatientID, _ = primitive.ObjectIDFromHex(c.GetString("userID"))
	appointment.ID = primitive.NewObjectID()

	collection := database.Database.Collection("appointments")
	_, err := collection.InsertOne(context.TODO(), appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment scheduled successfully"})
}

func ViewAvailableDoctors(c *gin.Context) {
	var doctors []models.Doctor
	collection := database.Database.Collection("doctors")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var doctor models.Doctor
		if err = cursor.Decode(&doctor); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		doctors = append(doctors, doctor)
	}

	if err = cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

func View_Appointments(c *gin.Context) {
	var appointments []models.Appointment
	patientId := c.GetString("userID")
	patientObjectId, _ := primitive.ObjectIDFromHex(patientId)

	collection := database.Database.Collection("appointments")
	cursor, err := collection.Find(context.TODO(), bson.M{"patient_id": patientObjectId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var appointment models.Appointment
		if err = cursor.Decode(&appointment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		appointments = append(appointments, appointment)
	}

	if err = cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}
