package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	/* create a slice */
	numbers := []int{0, 1, 7, 3, 9, 2, 6, 11, 2}

	/* print the numbers */
	for i, e := range numbers {
		fmt.Println("Slice item", i, "is", e)
	}

	/* print elements of a string */
	for s, c := range "John" {
		fmt.Printf("s: %v, d %v; ", s, c)
	}
	fmt.Println()

	/* create a map*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}

	fmt.Println("Done")
}
