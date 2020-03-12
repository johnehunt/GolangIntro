package main

import "fmt"

func main() {
	var greeting = "Hello world!"

	fmt.Printf("normal string: ")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")

	var sampleText = "John was here"
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q \n", sampleText)

	fmt.Printf("String Length is: ")
	fmt.Println(len(greeting))

	var s1 = "hello"
	var s2 = " world"
	var s3 = s1 + s2
	fmt.Printf("%s", s3)

}
