package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User representation - need to export
// the fields by capitalizing the first
// letter of the field name
// then providing a json mapping
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// users slice holds a list of users
var users = []User{}

// Set up some initial data
func init() {
	fmt.Println("Setting up initial users")
	users = append(users, User{234, "Denise", "Ireland"})
	users = append(users, User{987, "Phoebe", "USA"})
}

// Handle unknown URLs
func notFound(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte(`{"message": "not found"}`))
}

func getAllUsers(resp http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllUsers")
	resp.Header().Set("Content-Type", "application/json")

	usersJSON, err := json.Marshal(users)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(usersJSON)
}

func getUserDetails(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("getUserDetails")

	// Handle path parameters
	pathParams := mux.Vars(req)
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

	fmt.Printf("Building User object for %d, %s\n", userID, location)

	user := User{userID, "John", location}

	fmt.Println("user:", user)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(resp).Encode(&user); err != nil {
		fmt.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"message": "problem generating user info"}`))
	}

}

func postNewUser(resp http.ResponseWriter, req *http.Request) {

	fmt.Println("postNewUser")

	// Process post body
	var user User
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&user); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"message": "invalid payload"}`))
		return
	}
	defer req.Body.Close()

	// Add to in memory list of users
	users = append(users, user)

	// Generate response
	userJSON, _ := json.Marshal(&user)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	resp.Write(userJSON)
}

func putUser(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("postUser")

	// Process post body
	var user User
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&user); err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(`{"message": "invalid payload"}`))
		return
	}
	defer req.Body.Close()

	fmt.Println("Updating User", user)

	for i, u := range users {
		if u.ID == user.ID {
			users = append(users[:i], users[i+1:]...)
			users = append(users, user)
			break
		}
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message": "user updated"}`))

}

func deleteUser(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("deleteUser")

	// Handle path parameters
	pathParams := mux.Vars(req)
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

	fmt.Println("deleting User", userID)

	for i, u := range users {
		if u.ID == userID {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message": "removed user"}`))
}

func main() {
	fmt.Println("Starting Server")
	const urlPath = "/api"
	router := mux.NewRouter()
	// Create a subrouter - so don;t have to repeat path info
	api := router.PathPrefix(urlPath).Subrouter()
	// Handle GET requests
	api.HandleFunc("/users", getAllUsers).Methods(http.MethodGet)
	api.HandleFunc("/user/{userID}", getUserDetails).Methods(http.MethodGet)

	// Handle POST and PUT
	api.HandleFunc("/users", postNewUser).Methods(http.MethodPost)
	api.HandleFunc("/users", putUser).Methods(http.MethodPut)

	// Handle Delete
	api.HandleFunc("/user/{userID}", deleteUser).Methods(http.MethodDelete)

	api.HandleFunc("", notFound)
	fmt.Println("Server Available - see http://localhost:8080/api/user/123?location=UK")
	err := http.ListenAndServe(":8080", api)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
