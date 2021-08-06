package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people = []Person{{Name: "John", Age: 57}}

// Note this is a function not a method now
func serve(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case "GET":
		fmt.Println("processing GET Request")
		// To get a url parameter
		// e.g. http://localhost:8080/api/people/list?keys=hello
		keys, ok := req.URL.Query()["key"]

		if !ok {
			fmt.Println("Url Param 'key' is missing")
			return
		}

		// Query()["key"] will return an array of items,
		// we only want the single item.
		key := keys[0]

		fmt.Println("Url Param 'key' is: " + string(key))

		// Find path parameter
		id := strings.TrimPrefix(req.URL.Path, "/api/people/")
		fmt.Println("id:", id)

		if id == "" {
			// Want a list of all people
			jsonData, err := json.Marshal(people)
			if err != nil {
				fmt.Println("Something went wrong converting people to JSON!", err)
				msg := "Error"
				http.Error(resp, msg, http.StatusInternalServerError)
				return
			} else {
				resp.WriteHeader(http.StatusOK)
				resp.Write(jsonData)
			}
		} else {
			// return a single user
			jsonData, err := json.Marshal(people[0])
			if err != nil {
				fmt.Println("Something went wrong converting to JSON!", err)
				msg := "Error"
				http.Error(resp, msg, http.StatusInternalServerError)
				return
			} else {
				resp.WriteHeader(http.StatusOK)
				resp.Write(jsonData)
			}
		}
	case "POST":
		fmt.Println("processing POST Request")
		// Check the content header is appropriate
		contentheader := req.Header.Get("Content-Type")
		fmt.Println(contentheader)
		if contentheader != "" {
			if contentheader != "application/json" {
				msg := "Content-Type header is not application/json"
				http.Error(resp, msg, http.StatusUnsupportedMediaType)
				return
			}
		}
		// Now make sure body isn;t too large
		// Use http.MaxBytesReader to enforce a maximum read of 1MB from the
		// request body. A request body larger than that will now result in
		// Decode() returning a "http: request body too large" error.
		requestBody := http.MaxBytesReader(resp, req.Body, 1048576)

		// Now lets parse the body
		dec := json.NewDecoder(requestBody)
		dec.DisallowUnknownFields()

		var p Person
		err := dec.Decode(&p)
		if err != nil {
			fmt.Println("A problem with decoding the POST body")
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("Person:-", p)
		people = append(people, p)
		resp.WriteHeader(http.StatusCreated)
		fmt.Fprintf(resp, "Person: created")
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
	http.HandleFunc("/api/people/", serve)
	fmt.Println("Server Available - see http://localhost:8080/api/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error srarting server - ", err)
	}
}
