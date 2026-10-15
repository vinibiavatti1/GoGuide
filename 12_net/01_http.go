// HTTP (Hypertext Transfer Protocol)
// This file demonstrates the basic usage of HTTP servers and clients in Go, using the built-in net/http package.
// Ideal for REST APIs, web services, and client-side HTTP integrations.

package net

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Server
// Demonstrates a minimal HTTP server that listens on port 8081 and responds to standard GET requests.
func StartHTTPServer() {
	// Registering standard handler for the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from Go HTTP Server!")
	})

	fmt.Println("HTTP Server listening on :8081...")
	// Blocks and listens on port 8081
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

// Client
// Demonstrates an HTTP client making a GET request, reading the body, and closing the connection stream.
func StartHTTPClient() {
	resp, err := http.Get("http://localhost:8081/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // Essential to prevent resource leaks

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Client received:", string(body))
}

// TestHTTP
// Coordinates running both the Server and Client using Go's concurrency model (goroutines).
func TestHTTP() {
	go StartHTTPServer()
	time.Sleep(500 * time.Millisecond) // Let the server bind to the port
	StartHTTPClient()
	// Output:
	// HTTP Server listening on :8081...
	// Client received: Hello from Go HTTP Server!
}
