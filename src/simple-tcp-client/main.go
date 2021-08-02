package main

import (
	"fmt"
	"simple-tcp-client/client"
)

func main() {
	fmt.Println("Starting client app")

	client.CallServer("test server", "127.0.0.1", "3333", "tcp")

	fmt.Println("Done")
}
