// Panic
// The "panic" function is used to raise a panic in Go.
// A panic is a runtime error that occurs when the program encounters an unexpected condition.
// When a panic occurs, the program stops executing and starts to unwind the stack,
// looking for a deferred function that can handle the panic.
// The "recover" function is used to regain control of a panicking goroutine.

package errors

import "fmt"

// Declaring a Function that Raises a Panic
// The function below raises a panic.
// Usually, only critical errors are raised as panics.
// In this case, we are simulating a configuration error.
func Configuration() {
	panic("Configuration error!")
}

// Unhandled Panic
// The function below raises a panic and does not handle it.
// This will cause the program to terminate immediately.
func UnhandledPanic() {
	Configuration()
	// Output: panic: Configuration error!
}

// Handled Panic
// The unique way to handle a panic is to use the "defer" statement in the function that raises the panic.
// The "recover" function is used to regain control of a panicking goroutine.
func HandledPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r) // Recovered from panic: Configuration error!
		}
	}()
	Configuration() // Raises a panic
}
