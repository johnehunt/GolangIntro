package main

import (
	"fmt"
	"strings"
)

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

	// true
	fmt.Println(strings.Contains("test", "es"))

	// 2
	fmt.Println(strings.Count("test", "t"))

	// true
	fmt.Println(strings.HasPrefix("test", "te"))

	// true
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	// == "aaaaa"
	fmt.Println(strings.Repeat("a", 5))

	// "bbaa"
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))

	// []string{"a","b","c","d","e"}
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	fmt.Println(strings.ToLower("John"))
	fmt.Println(strings.ToUpper("john"))

}
