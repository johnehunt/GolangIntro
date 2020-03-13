package calc

import "fmt"

// InvalidArgumentError custom exception type
// Implements the Error interface with the Errors method
type InvalidArgumentError struct {
	argument int
	problem  string
}

func (e *InvalidArgumentError) Error() string {
	return fmt.Sprintf("%d - %s", e.argument, e.problem)
}

func calculator(operation string, x int, y int) (int, error) {
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
