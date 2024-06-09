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

func UpdatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.BindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patientId, err := primitive.ObjectIDFromHex(c.Param("patientId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	collection := database.Database.Collection("patients")
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": patientId},
		bson.M{"$set": patient},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient updated successfully"})
}

func DeletePatient(c *gin.Context) {
	patientId, err := primitive.ObjectIDFromHex(c.Param("patientId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	collection := database.Database.Collection("patients")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": patientId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
}

func ListPatients(c *gin.Context) {
	var patients []models.Patient
	collection := database.Database.Collection("patients")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var patient models.Patient
		if err = cursor.Decode(&patient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		patients = append(patients, patient)
	}

	if err = cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func ViewAppointments(c *gin.Context) {
	var appointments []models.Appointment
	doctorId := c.GetString("userID")
	doctorObjectId, _ := primitive.ObjectIDFromHex(doctorId)

	collection := database.Database.Collection("appointments")
	cursor, err := collection.Find(context.TODO(), bson.M{"doctor_id": doctorObjectId})
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

func SetAvailability(c *gin.Context) {
	var availability models.Availability
	if err := c.BindJSON(&availability); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	availability.DoctorID, _ = primitive.ObjectIDFromHex(c.GetString("userID"))

	collection := database.Database.Collection("availability")
	availability.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), availability)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Availability set successfully"})
}
