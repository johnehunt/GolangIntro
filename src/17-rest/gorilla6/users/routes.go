package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GetRoutes Public method to generate routes
func GetRoutes() *mux.Router {
	const urlPath = "/api"
	router := mux.NewRouter()
	// Create a subrouter - so don;t have to repeat path info
	api := router.PathPrefix(urlPath).Subrouter()
	// Handle GET requests
	api.HandleFunc("/users", GetAllUsers).Methods(http.MethodGet)
	api.HandleFunc("/user/{userID}", GetUserDetails).Methods(http.MethodGet)

	// Handle POST and PUT
	api.HandleFunc("/users", PostNewUser).Methods(http.MethodPost)
	api.HandleFunc("/users", PutUpdateUser).Methods(http.MethodPut)

	// Handle Delete
	api.HandleFunc("/user/{userID}", DeleteUser).Methods(http.MethodDelete)

	api.HandleFunc("", NotFound)

	return api

}
