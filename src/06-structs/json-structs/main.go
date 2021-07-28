package main

import (
	"encoding/json"
	"fmt"
)

// Book a book structure
type Book struct {
	ISBN   int
	Title  string
	Author string  `json:"authorid"`
	Price  float64 `json:"price"`
}

func main() {

	var books = []Book{
		Book{ISBN: 123, Title: "Go Programming", Author: "a1", Price: 12.99},
		Book{ISBN: 322, Title: "Python Programming", Author: "a2", Price: 11.55},
	}
	fmt.Printf("books: %v\n", books)

	// Marshal to JSON
	book := Book{123, "Go Introduction", "Bill", 12.99}
	fmt.Println("Book:", book)
	data, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Something went wrong converting to JSON!", err)
	} else {
		fmt.Printf("Data: %s\n", data)
	}

	fmt.Println("-------------")

	data, err = json.Marshal(books)
	if err != nil {
		fmt.Println("Something went wrong un marshalling from JSON!", err)
	} else {
		fmt.Printf("Data: %s\n", data)
	}

	// Unmarshal JSON - populate slice with book data

	var newbooks []Book
	err2 := json.Unmarshal(data, &newbooks)
	if err2 != nil {
		fmt.Println("Something went wrong!", err2)
	} else {
		fmt.Printf("newbooks: %v\n", newbooks)
	}

}
