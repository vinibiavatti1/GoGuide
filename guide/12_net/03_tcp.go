// TCP/IP (Transmission Control Protocol/Internet Protocol) Sockets
// This file demonstrates streaming connection-oriented TCP socket communication in Go.
// TCP Sockets guarantee reliable, ordered, and error-checked delivery of a stream of bytes.

package net

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// Server
// Demonstrates a structured TCP server listening on port 8083, accepting connections, and parsing text via bufio.
func StartTCPServer() {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("TCP Server listening on :8083...")

	conn, err := listener.Accept() // Blocks waiting for client connection handshake
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Receive Text via Buffered Reader
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n') // Blocks and reads until the newline delimiter
	if err != nil {
		panic(err)
	}
	fmt.Print("Server received: ", message)

	// Send Response back using raw string converted to bytes
	_, err = conn.Write([]byte("Hi Client!\n"))
	if err != nil {
		panic(err)
	}
}

// Client
// Demonstrates a TCP client connecting to localhost, pushing a message string, and reading the synchronous server feedback.
func StartTCPClient() {
	conn, err := net.Dial("tcp", "localhost:8083")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send Message
	_, err = conn.Write([]byte("Hi Server!\n"))
	if err != nil {
		panic(err)
	}

	// Receive Response
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Print("Client received: ", response)
}

// TestTCP
// Orchestrates the client-server interaction over streaming TCP sockets.
func TestTCP() {
	go StartTCPServer()
	time.Sleep(500 * time.Millisecond) // Ensures server socket bound completely
	StartTCPClient()
	// Output:
	// TCP Server listening on :8083...
	// Server received: Hi Server!
	// Client received: Hi Client!
}
