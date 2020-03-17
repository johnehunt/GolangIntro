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

	// Find One -------------------------

	var query = bson.D{{"id", 123}}
	// create a value into which the result can be decoded
	var result User
	err = collection.FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found a single document: %+v\n", result)

	// Find Many ------------------------

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(100)

	// Here's an array in which you can store the decoded documents
	var results []*User

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var user User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("user:", user)
		results = append(results, &user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	fmt.Println("Disconnecting from MondoDB")

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

	fmt.Println("Done")
}
