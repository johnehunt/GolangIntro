package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting HTTP CLient")
	resp, err := http.Get("http://localhost:8080/api/people/1")
	if err != nil {
		fmt.Println("Problem connecting to server", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("Response status:", resp.Status)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Problem reading body", err)
		}
	}

}
