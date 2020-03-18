package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// Employee holds info on an employee
type Employee struct {
	ID   int
	name string
}

func main() {
	fmt.Println("Go MySQL Tutorial")

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

	const sql = "UPDATE employees.employee SET name = 'Jasmine Bowen' WHERE id = 987"

	// Run a query to insert a row into the database
	update, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Record updated")
	}
	defer update.Close()

}
