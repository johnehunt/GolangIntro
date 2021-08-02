package main

import (
	"fmt"
	"simple-tcp-client/client"
)

func main() {
	fmt.Println("Starting client app")

	client.CallServer("test server", "localhost", "3333", "tcp")

	fmt.Println("Done")
}
