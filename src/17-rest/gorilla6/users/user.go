package users

import (
	"fmt"
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

// Users slice holds a list of users
var Users = []User{}

// Set up some initial data
func init() {
	fmt.Println("Setting up initial users")
	Users = append(Users, User{234, "Denise", "Ireland"})
	Users = append(Users, User{987, "Phoebe", "USA"})
}
