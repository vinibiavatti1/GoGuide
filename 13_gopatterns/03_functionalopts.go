// Functional Options
// - Functional options are a common pattern in Go for configuring structs or functions.
// - They allow you to pass a variable number of options to a function, making it more flexible and
//   easier to read.
// - This pattern is often used in constructors or initialization functions.
// - The idea is to define a function type that takes a pointer to the struct you want to configure,
//   and then create a series of functions that implement this type.

package gopatterns

import "fmt"

//==================================================================================================
// Declare Functional Opts
//==================================================================================================

// Struct
// - The struct below will be used to demonstrate the functional options pattern.
// - We will define a single constructor function that takes a variable number of options to
//   configure the struct.

type Server struct {
	Host string
	Port int
	TLS  bool
}

// Option Functions
// - These functions return the Option type and allow us to set various fields in the Server struct.
// - Each function takes a pointer to the Server struct and modifies it accordingly.

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.Host = host
	}
}
func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.Port = port
	}
}
func WithTLS(enabled bool) func(*Server) {
	return func(s *Server) {
		s.TLS = enabled
	}
}

// Constructor
// - Note that the constructor below is more dynamic, since it can accept any number of options.
// - This allows us to create a Server instance with various configurations without needing to
//   define multiple constructors.

func NewServer(opts ...func(*Server)) *Server {
	s := &Server{
		Host: "localhost",
		Port: 8080,
		TLS:  false,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

//==================================================================================================
// Test Functional Opts
//==================================================================================================

func TestFunctionalOpts() {
	s1 := NewServer(
		WithHost("https://test.com"),
		WithPort(443),
		WithTLS(true),
	)
	s2 := NewServer(
		WithTLS(true),
	)
	fmt.Println(s1) // Output: &{https://test.com 443 true}
	fmt.Println(s2) // Output: &{localhost 8080 true}
}
