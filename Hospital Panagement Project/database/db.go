package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database

func ConnectDatabase() {
	// Set MongoDB connection options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Create a MongoDB client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Create a context with a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB database
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Select the database
	Database = client.Database("hospital")
	Client = client

	// Ping the database to ensure a successful connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}
