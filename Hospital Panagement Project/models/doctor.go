package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Doctor struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Specialty string             `bson:"specialty,omitempty"`
}
