package main

import (
	"17-rest/gorilla6/users"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Server")
	const urlPath = "/api"
	routes := users.GetRoutes()

	fmt.Println("Server Available - see http://localhost:8080/api/user/123?location=UK")
	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
