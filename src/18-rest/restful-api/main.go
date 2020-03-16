package main

import (
	"fmt"
	"net/http"
)

// Server Represents the server type
type Server struct{}

// ServerHTTP - implements the Handler interface
func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case "GET":
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(`{"message": "get called"}`))
	case "POST":
		resp.WriteHeader(http.StatusCreated)
		resp.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		resp.WriteHeader(http.StatusAccepted)
		resp.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte(`{"message": "delete called"}`))
	default:
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	fmt.Println("Starting Server")
	server := &Server{}
	// Set up a mapping from url to server
	http.Handle("/api/", server)
	fmt.Println("Server Available - see http://localhost:8080/api/")
	// Listen and server is a blocking call
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
