package main

import (
	"fmt"
	"simple-tcp-server/server"
)

func main() {
	fmt.Println("Starting Server")

	server.Start("127.0.0.1", "3333", "tcp")

	fmt.Println("Done")
}
