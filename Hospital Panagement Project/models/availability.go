package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Availability struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	DoctorID primitive.ObjectID `bson:"doctor_id,omitempty"`
	Date     string             `bson:"date,omitempty"`
	Slot     string             `bson:"slot,omitempty"`
}
