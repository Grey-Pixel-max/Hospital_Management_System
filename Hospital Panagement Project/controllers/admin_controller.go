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

func CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.BindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.Database.Collection("doctors")
	doctor.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), doctor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor created successfully"})
}

func UpdateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.BindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctorId, err := primitive.ObjectIDFromHex(c.Param("doctorId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	collection := database.Database.Collection("doctors")
	_, err = collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": doctorId},
		bson.M{"$set": doctor},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor updated successfully"})
}

func DeleteDoctor(c *gin.Context) {
	doctorId, err := primitive.ObjectIDFromHex(c.Param("doctorId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid doctor ID"})
		return
	}

	collection := database.Database.Collection("doctors")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": doctorId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}

func ListDoctors(c *gin.Context) {
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

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.BindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.Database.Collection("patients")
	patient.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient created successfully"})
}

func Update_Patient(c *gin.Context) {
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

func Delete_Patient(c *gin.Context) {
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

func List_Patients(c *gin.Context) {
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

func ViewAllppointments(c *gin.Context) {
	var appointments []models.Appointment
	collection := database.Database.Collection("appointments")

	cursor, err := collection.Find(context.TODO(), bson.M{})
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
