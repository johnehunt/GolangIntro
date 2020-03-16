package main

import (
	"fmt"
	"net/http"
	"strconv"

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

func getUserDetails(resp http.ResponseWriter, req *http.Request) {

	// Handle path parameters
	pathParams := mux.Vars(req)
	resp.Header().Set("Content-Type", "application/json")
	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte(`{"message": "need a userid"}`))
			return
		}
	}

	// Handle query params (if present)
	query := req.URL.Query()
	location := query.Get("location")

	// Generate repsonse info
	resp.Write([]byte(fmt.Sprintf(`{"userID": %d, "location": "%s" }`, userID, location)))
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

	// Example handling path parameters
	api.HandleFunc("/user/{userID}", getUserDetails).Methods(http.MethodGet)

	api.HandleFunc("", notFound)
	fmt.Println("Server Available - see http://localhost:8080/api/user/123?location=UK")
	err := http.ListenAndServe(":8080", api)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
