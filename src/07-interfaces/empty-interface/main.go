package main

import "fmt"

func describe(v interface{}) {
	// Check for type of data held in interface
	switch v.(type) {
	case int:
		fmt.Println("Its an int:", v.(int))
	case string:
		fmt.Println("Its an string:", v.(string))
	case float64:
		fmt.Println("Its an float64:", v.(float64))
	case bool:
		fmt.Println("Its an boolean:", v.(bool))
	default:
		fmt.Println("It is some other type")
	}
}

func main() {
	var v1 interface{}
	fmt.Println("Default value of an interface:", v1) // Its <nil>

	var v2 interface{} = "John"
	fmt.Println("v2:", v2) // Its "John"

	fmt.Println("v2.(string):", v2.(string))

	// OK is a flag to indicate whether value was of given type
	value, ok := v2.(float64)
	fmt.Printf("v2.(string): value %f, ok %t\n", value, ok)

	describe(v2)

}
