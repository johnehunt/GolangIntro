package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func get(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message": "get called"}`))
}

func post(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	resp.Write([]byte(`{"message": "post called"}`))
}

func put(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusAccepted)
	resp.Write([]byte(`{"message": "put called"}`))
}

func delete(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message": "delete called"}`))
}

func notFound(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte(`{"message": "not found"}`))
}

func main() {
	fmt.Println("Starting Server")
	const urlPath = "/api/"
	router := mux.NewRouter()
	// Create a subrouter - so don;t have to repeat path info
	api := router.PathPrefix(urlPath).Subrouter()
	api.HandleFunc("", get).Methods(http.MethodGet)
	api.HandleFunc("", post).Methods(http.MethodPost)
	api.HandleFunc("", put).Methods(http.MethodPut)
	api.HandleFunc("", delete).Methods(http.MethodDelete)
	api.HandleFunc("", notFound)
	fmt.Println("Server Available - see http://localhost:8080/api/")
	err := http.ListenAndServe(":8080", api)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
