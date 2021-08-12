package calc

import "errors"

const PLUS, SUB, MULT, DIV = "+", "-", "*", "/"

var InvalidArgumentError= errors.New("Invalid Argument")
var InvalidOperationError= errors.New("Invalid Argument")

func Calculator(operation string, x int, y int) (int, error) {
	var result = 0
	switch operation {
	case PLUS:
		result = x + y
	case SUB:
		result = x - y
	case MULT:
		result = x * y
	case DIV:
		if y == 0 {
			return -1, InvalidArgumentError
		}
		result = x / y
	default:
		return -1, InvalidOperationError
	}
	return result, nil
}

