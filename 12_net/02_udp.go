// UDP (User Datagram Protocol) Sockets
// This file demonstrates the basic usage of connectionless UDP communication in Go.
// UDP is a connectionless, unreliable transport layer protocol.
// Packets (datagrams) may arrive out of order or be dropped entirely, but it has very low overhead.

package net

import (
	"fmt"
	"net"
	"time"
)

// Server
// Demonstrates a UDP server that binds to a port, listens for raw datagrams, and reads them into a buffer.
func StartUDPServer() {
	addr, err := net.ResolveUDPAddr("udp", ":8082")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("UDP Server listening on :8082...")

	buffer := make([]byte, 1024)
	// ReadFromUDP reads a packet and returns the number of bytes and the client's remote address
	n, remoteAddr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Received from UDP Client (%s): %s\n", remoteAddr, string(buffer[:n]))
}

// Client
// Demonstrates a UDP client dialing a remote endpoint and instantly writing a payload over a datagram packet.
func StartUDPClient() {
	remoteAddr, err := net.ResolveUDPAddr("udp", "localhost:8082")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := []byte("Ping over UDP!")
	_, err = conn.Write(message)
	if err != nil {
		panic(err)
	}
}

// TestUDP
// Drives the execution of the independent UDP server and client.
func TestUDP() {
	go StartUDPServer()
	time.Sleep(500 * time.Millisecond)
	StartUDPClient()
	// Output:
	// UDP Server listening on :8082...
	// Received from UDP Client (127.0.0.1:XXXXX): Ping over UDP!
}
