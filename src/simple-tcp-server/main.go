package main

import (
	"fmt"
	"simple-tcp-server/server"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	fmt.Println("Starting Server")

	server.Start(CONN_HOST, CONN_PORT, CONN_TYPE)

	fmt.Println("Done")
}
