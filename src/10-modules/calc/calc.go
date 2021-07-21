package calc

// make sure you run go install 10-modules/calc to compile
// the calculcator.go file and create a linkable object in the package location

import "fmt"

func init() {
	fmt.Println("Runs when package calc is loaded")
}

// InvalidArgumentError custom exception type
// Implements the Error interface with the Errors method
type InvalidArgumentError struct {
	argument int
	problem  string
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("%d - %s", e.argument, e.problem)
}

// Calculator Function to perform calculations
func Calculator(operation string, x int, y int) (int, error) {
	var result = 0
	switch operation {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			return -1, &InvalidArgumentError{y, "No division by Zero"}
		}
		result = x / y
	default:
		return -1, &InvalidArgumentError{y, "Unknown operation"}
	}
	return result, nil
}
