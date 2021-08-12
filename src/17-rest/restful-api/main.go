package main

import (
	"fmt"
	"net/http"
)

// Server Represents the server type
type Server struct{}

// ServerHTTP - implements the Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case http.MethodPut:
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case http.MethodDelete:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
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
