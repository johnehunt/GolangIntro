package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// NotFound Handle unknown URLs
func NotFound(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusNotFound)
	resp.Write([]byte(`{"message": "not found"}`))
}

// GetAllUsers return all known users
func GetAllUsers(resp http.ResponseWriter, r *http.Request) {
	fmt.Println("getAllUsers")
	resp.Header().Set("Content-Type", "application/json")

	usersJSON, err := json.Marshal(Users)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(usersJSON)
}

// GetUserDetails for a specified userID
func GetUserDetails(resp http.ResponseWriter, req *http.Request) {
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

// PostNewUser add a new user based on request body JSON data
func PostNewUser(resp http.ResponseWriter, req *http.Request) {

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
	Users = append(Users, user)

	// Generate response
	userJSON, _ := json.Marshal(&user)
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusCreated)
	resp.Write(userJSON)
}

// PutUpdateUser update a user based on request body JSON data
func PutUpdateUser(resp http.ResponseWriter, req *http.Request) {
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

	for i, u := range Users {
		if u.ID == user.ID {
			Users = append(Users[:i], Users[i+1:]...)
			Users = append(Users, user)
			break
		}
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message": "user updated"}`))

}

// DeleteUser based on userID
func DeleteUser(resp http.ResponseWriter, req *http.Request) {
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

	for i, u := range Users {
		if u.ID == userID {
			Users = append(Users[:i], Users[i+1:]...)
			break
		}
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"message": "removed user"}`))
}
