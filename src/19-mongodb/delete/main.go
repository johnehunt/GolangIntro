package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

	// Delete One --------------
	query := bson.D{{"id", 111}}
	deleteResult, err := collection.DeleteOne(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the users collection\n", deleteResult.DeletedCount)

	// Delete Many -------------
	query = bson.D{{"name", "John Smith"}}
	deleteResult, err = collection.DeleteMany(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the users collection\n", deleteResult.DeletedCount)

	fmt.Println("Disconnecting from MondoDB")

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

	fmt.Println("Done")
}
