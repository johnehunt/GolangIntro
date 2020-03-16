package main

import (
	"fmt"
	"net/http"
)

// Server Represents the server type
type Server struct{}

// ServerHTTP - implements the Handler interface
func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	fmt.Println("Starting Server")
	server := &Server{}
	// Set up a mapping from url to server
	http.Handle("/", server)
	fmt.Println("Server Available - see http://localhost:8080/")
	// Listen and server is a blocking call
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
