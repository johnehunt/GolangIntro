package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting")

	// Crates a variable that can hold a map but initiliaed to zeroth value - nil
	var countryCapitalMap map[string]string
	fmt.Println(countryCapitalMap)

	/* instantiate a map */
	countryCapitalMap = make(map[string]string)

	/* insert key-value pairs in the map*/
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	fmt.Println("------------")

	/* test if entry is present in the map or not*/
	capital, ok := countryCapitalMap["United States"]
	/* if ok is true, entry is present otherwise entry is absent*/
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	/* delete an entry */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	/* print map using key-value*/
	for country, capital := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", capital)
	}

	fmt.Println("------------")

	fmt.Println("Done")
}
