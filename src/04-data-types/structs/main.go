package main

import (
	"fmt"
)

// Book structure
type Book struct {
	title   string
	author  string
	subject string
	isbn    int
}

func main() {
	fmt.Println("Starting")

	var Book1 Book /* Declare Book1 of type Book */
	var Book2 Book /* Declare Book2 of type Book */

	/* book 1 specification */
	Book1.title = "Go is the Way"
	Book1.author = "John Hunt"
	Book1.subject = "Go Programming"
	Book1.isbn = 6495407

	/* book 2 specification */
	Book2.title = "Python Introduction"
	Book2.author = "Denise Jonesi"
	Book2.subject = "Python Tutorial"
	Book2.isbn = 6495700

	/* print Book1 info */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.isbn)

	/* print Book2 info */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.isbn)

	// Pointers

	var structPointer *Book
	structPointer = &Book1
	fmt.Println("structPointer.title:", structPointer.title)

	fmt.Println("Done")
}
