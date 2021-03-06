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

	// Run a query on the database - returns a Rows
	employees, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err.Error())
	}

	// Loop through the results returned
	for employees.Next() {
		// For each row convert into a struct
		var employee Employee
		err := employees.Scan(&employee.ID, &employee.name)
		if err != nil {
			panic(err.Error())
		}
		// Print the employee details
		fmt.Println(employee)
	}

	fmt.Println("Defer closing to end of function")
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

}
