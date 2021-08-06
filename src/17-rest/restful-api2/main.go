package main

import (
	"fmt"
	"net/http"
)

// Note this is a function not a method now
func serve(resp http.ResponseWriter, req *http.Request) {
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
	// HandleFunc is a method on http that has the same signature as
	// ServeHTTP and can serve a route - avoids need for struct
	http.HandleFunc("/api/", serve)
	fmt.Println("Server Available - see http://localhost:8080/api/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
