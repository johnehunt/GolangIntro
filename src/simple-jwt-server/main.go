package main

import (
	"fmt"
	"net/http"
	"simple-jwt-server/server"
)

func main() {
	fmt.Println("Starting Server App")

	fmt.Println("Registering functions with urls")
	http.HandleFunc("/signin", server.Signin)
	http.HandleFunc("/welcome", server.Welcome)

	// Start the server on port 8000
	fmt.Println("Starting Listen and Serve")
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		fmt.Printf("Error starting server %v", err)
	}
}
