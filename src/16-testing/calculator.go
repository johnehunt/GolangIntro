package calc

import (
	"errors"
)

// Public constants
const ADD, SUB, MULT, DIV = "+", "-", "*", "/"

// InvalidArgumentError custom exception type
// Implements the Error interface with the Errors method
var ErrorInvalidArgument = errors.New("invalid argument")
var ErrorInvalidOperation = errors.New("invalid argument")

// Calculator function
func Calculator(operation string, x int, y int) (int, error) {
	var result = 0
	switch operation {
	case ADD:
		result = x + y
	case SUB:
		result = x - y
	case MULT:
		result = x * y
	case DIV:
		if y == 0 {
			return -1, ErrorInvalidArgument
		}
		result = x / y
	default:
		return -1, ErrorInvalidOperation
	}
	return result, nil
}
