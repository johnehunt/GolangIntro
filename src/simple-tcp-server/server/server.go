package server

import (
	"fmt"
	"net"
	"os"
	"time"
)

func handleRequest(conn net.Conn) {
	fmt.Println("Calling handleRequest")

	// Create a slice of type byte
	// Will be used to hold incoming data
	buffer := make([]byte, 1024)
	defer conn.Close()

	// Read the incoming data into the budeer
	_, err := conn.Read(buffer)
	if err != nil {
		// Somethign went wrong
		fmt.Printf("Error: %v", err)
		conn.Write([]byte("Message received with an error."))
	} else {
		currentTime := time.Now()
		// All ok so handle response
		fmt.Println("Message recieved as: ", string(buffer))
		fmt.Println("Generating response")
		conn.Write([]byte("Message received with "))
		conn.Write(buffer)
		conn.Write([]byte(" at " + currentTime.String() + "\n"))
	}

}

// Start call this function to start the server listening
// on a given host, port using a connection type
//
// Can test using:
//      echo -n "test message" | nc 127.0.0.1 3333
func Start(host string, port string, connectionType string) {
	fmt.Println("Starting server")

	listener, err := net.Listen(connectionType, host+":"+port)
	if err != nil {
		fmt.Printf("Error openning server: %v", err)
	}
	defer listener.Close()

	fmt.Printf("Listening on %s:%s using %s\n", host, port, connectionType)
	// Start listening loop
	for {
		// Accept a request
		conn, err := listener.Accept()
		if err != nil {
			// SOmething went wrong - log it and terminate
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Otherwise handle the request in a GoRoutine
		go handleRequest(conn)
	}
}
