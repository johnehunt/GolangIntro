package main

import (
	"encoding/json"
	"fmt"
)

// Book a book structure
type Book struct {
	ISBN   int
	Title  string
	Author string `json:"authorid"`
	Price  float64
}

func main() {

	var books = []Book{
		Book{ISBN: 123, Title: "Go Programming", Author: "a1", Price: 12.99},
		Book{ISBN: 322, Title: "Python Programming", Author: "a2", Price: 11.55},
	}
	fmt.Printf("books: %v\n", books)

	// Marshal to JSON
	data, err := json.Marshal(books[0])
	if err != nil {
		fmt.Println("Something went wrong!", err)
	} else {
		fmt.Printf("Data: %s\n", data)
	}

	data, err = json.Marshal(books)
	if err != nil {
		fmt.Println("Something went wrong!", err)
	} else {
		fmt.Printf("Data: %s\n", data)
	}

	// Unmarshal JSON

	var newbooks [2]Book
	err2 := json.Unmarshal(data, &newbooks)
	if err2 != nil {
		fmt.Println("Something went wrong!", err2)
	} else {
		fmt.Printf("newbooks: %v\n", newbooks)
	}

}
