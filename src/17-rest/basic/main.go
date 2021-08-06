package main

import (
	"fmt"
	"net/http"
)

// Server Represents the server type
type Server struct{}

// ServerHTTP - implements the Handler interface
func (s *Server) ServeHTTP(resp http.ResponseWriter, req *http.Request) {

	fmt.Printf("req.Method: %s \n", req.Method)
	fmt.Printf("req.URL: %s \n", req.URL)
	fmt.Printf("req.Header: %s \n", req.Header)
	fmt.Printf("req.Body: %s \n", req.Body)

	resp.WriteHeader(http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	fmt.Println("Starting Server")
	server := &Server{}
	// Set up a mapping from url to server
	http.Handle("/api/", server)

	fmt.Println("Server Available - see http://localhost:8080/api/")
	// Listen and serve is a blocking call
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server - ", err)
	}
}
