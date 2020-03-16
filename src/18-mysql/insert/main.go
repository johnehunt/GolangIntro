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

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/employees")

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("Problem Encountered", err)
		panic(err.Error())
	}

	fmt.Println("Defer closing to end of function")
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	const sql = "INSERT INTO employees.employee (id, name) VALUES (987, 'Jas Bown')"

	// Run a query to insert a row into the database
	insert, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

}
