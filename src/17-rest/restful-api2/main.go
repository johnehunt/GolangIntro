package main

import (
	"fmt"
	"net/http"
)

// Note this is a function not a method now
func serve(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
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
	// server is a plain old function
	http.HandleFunc("/api/", serve)
	fmt.Println("Server Available - see http://localhost:8080/api/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
