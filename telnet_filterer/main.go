package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var clientFrequency = make(map[string]int)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Get the client's IP address
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Connection from %s\n", clientAddr)

	// Send a welcome message
	conn.Write([]byte("Welcome to the Floaty Telnet Filter!\n"))

	// Create a bufio scanner to read input from the client
	scanner := bufio.NewScanner(conn)

	clientFrequency[clientAddr] = 0
	go func() {
		// reset clientFrequency every 5 seconds
		for {
			time.Sleep(5 * time.Second)
			clientFrequency[clientAddr] = 0
		}
	}()
	// Process incoming messages
	for scanner.Scan() {
		message := scanner.Text()
		if clientFrequency[clientAddr] >= 3 {
			conn.Write([]byte("Message frequency too high. Please wait a little bit.\n"))
			continue
		}

		clientFrequency[clientAddr]++
		// Print the received message
		fmt.Printf("%s\n", message)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Telnet server started on port 5555")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection concurrently
		go handleConnection(conn)
	}
}
