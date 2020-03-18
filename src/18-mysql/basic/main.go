package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Starting")

	fmt.Println("Open up database connection")
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/employees")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Problem Encountered", err)
		panic(err.Error())
	} else {
		fmt.Println("Successfully opened database connection")
	}

	fmt.Println("Defer closing to end of function")
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	fmt.Println("Done")
}
