package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Appointment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	DoctorID  primitive.ObjectID `bson:"doctor_id,omitempty"`
	PatientID primitive.ObjectID `bson:"patient_id,omitempty"`
	Date      string             `bson:"date,omitempty"`
}
