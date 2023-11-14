package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	// Create a new Scanner to read input from the client
	scanner := bufio.NewScanner(conn)

	// Read input from the client in a loop
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Println("Received:", input)

		// Send back response
		response := []byte("Command FAILED\r\n")
		conn.Write(response)
		fmt.Printf("Sent response %s\n", response)
	}

	fmt.Println("Client disconnected:", conn.RemoteAddr())
}

func main() {
	// Set the port to listen on
	port := ":12345"

	// Listen for incoming connections
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Telnet server listening on", port)

	// Accept and handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}
