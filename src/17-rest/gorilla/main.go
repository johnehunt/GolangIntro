package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func get(resp http.ResponseWriter, r *http.Request) {
	fmt.Println("In get request")
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
	router.HandleFunc(urlPath, get).Methods(http.MethodGet)
	router.HandleFunc(urlPath, post).Methods(http.MethodPost)
	router.HandleFunc(urlPath, put).Methods(http.MethodPut)
	router.HandleFunc(urlPath, delete).Methods(http.MethodDelete)
	router.HandleFunc(urlPath, notFound)
	fmt.Println("Server Available - see http://localhost:8080/api/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
