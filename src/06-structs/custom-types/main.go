package main

import "fmt"

type ResponseCode int

func main() {
	Handle(200)
}

func Handle(code ResponseCode) {
	switch code {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("Not Found")
	case 418:
		fmt.Println("I’m a Teapot")
	case 500:
		fmt.Println("Server Error")
	default:
		fmt.Printf("Unknown code %d\n", code)
	}
	fmt.Printf("%T\n", code)
}
