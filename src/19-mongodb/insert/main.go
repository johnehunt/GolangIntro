package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User holds info on an employee
type User struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("Starting")

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongo Database!")

	fmt.Println("Connecting to collection")
	collection := client.Database("userdb").Collection("users")

	// Insert One --------------
	newUser := User{454, "Steve Jones"}
	insertResult, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert many --------------
	newUsers := []interface{}{
		User{333, "Pete Best"},
		User{433, "David Anderson"}}

	insertManyResult, err := collection.InsertMany(context.TODO(), newUsers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	fmt.Println("Disconnecting from MondoDB")

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

	fmt.Println("Done")
}
