package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// CallServer - send a message to a server at the given
// server, port and connectionType
func CallServer(msg string, server string, port string, connectionType string) {
	fmt.Printf("Calling server %s:%s using %s\n", server, port, connectionType)
	// connect to server
	conn, err := net.Dial(connectionType, server+":"+port)

	if err != nil {
		// Something went wrong
		fmt.Printf("Error connecting to server: %v", err)
	} else {
		fmt.Println("Connection successful")
		defer conn.Close()
		fmt.Println("Sending test message")
		if _, err := conn.Write([]byte("test message")); err != nil {
			fmt.Printf("Error writing to server: %v", err)
		} else {
			// Now wait for response
			response, err := bufio.NewReader(conn).ReadString('\n')
			switch err {
			case nil:
				fmt.Print("Received message from server:\n\t")
				fmt.Println(response)
			case io.EOF:
				fmt.Printf("Server closed the connection")
			default:
				fmt.Printf("Error reading from server: %v", err)
			}
		}
	}
}
