package main

import (
	"fmt"
)

// Book structure
type Book struct {
	Title   string
	Author  string
	Subject string
	ISBN    int
}

func main() {
	fmt.Println("Starting")

	var book1 Book /* Declare Book1 of type Book */
	var book2 Book /* Declare Book2 of type Book */

	/* book 1 specification */
	book1.Title = "Go is the Way"
	book1.Author = "John Hunt"
	book1.Subject = "Go Programming"
	book1.ISBN = 6495407

	/* book 2 specification */
	book2.Title = "Python Introduction"
	book2.Author = "Denise Jonesi"
	book2.Subject = "Python Tutorial"
	book2.ISBN = 6495700
	fmt.Printf("book3: %v\n", book2)

	/* print Book1 info */
	fmt.Printf("Book 1 Title : %s\n", book1.Title)
	fmt.Printf("Book 1 Author : %s\n", book1.Author)
	fmt.Printf("Book 1 Subject : %s\n", book1.Subject)
	fmt.Printf("Book 1 book_id : %d\n", book1.ISBN)

	/* print Book2 info */
	fmt.Printf("Book 2 Title : %s\n", book2.Title)
	fmt.Printf("Book 2 Author : %s\n", book2.Author)
	fmt.Printf("Book 2 Subject : %s\n", book2.Subject)
	fmt.Printf("Book 2 book_id : %d\n", book2.ISBN)

	// Struct Literal
	var book3 = Book{"JavaScript Uncovered", "Phoebe", "Technical", 1234}
	fmt.Printf("book3: %v\n", book3)

	// Pointers

	var structPointer *Book
	structPointer = &book1
	fmt.Println("structPointer.Title:", structPointer.Title)

	fmt.Println("Done")
}
