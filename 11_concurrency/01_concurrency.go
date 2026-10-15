// Concurrency
// - Concurrency is the ability of a program to make progress on multiple tasks at the same time.
// - In Go, concurrency is achieved through goroutines and channels.
// - Goroutines are lightweight threads managed by the Go runtime. They are created using the "go"
//   keyword followed by a statement.
// - Channels are used to communicate between goroutines. They can be thought of as pipes that
//   connect goroutines.
// - If a channel is unbuffered, the sending goroutine will block until the receiving goroutine is
//   ready to receive the value.
// - If a channel is buffered, the sending goroutine will block only if the buffer is full.
// - Syntax:
//   - chan <type>    : Bidirectional channel
//   - <-chan <type>  : Receive-only channel
//   - chan<- <type>  : Send-only channel
//   - go <statement> : Start a new goroutine
//   - <-ch           : Receive from channel ch
//   - ch <- <value>  : Send value to channel ch

package concurrency

import (
	"fmt"
	"time"
)

//==================================================================================================
// Declaring Channels
//==================================================================================================

// Declaring a Bidirectional Channel
// - A bidirectional channel can be used to send and receive values of a specific type.

var _ chan int

// Declaring a Send-Only Channel
// - A send-only channel can only be used to send values of a specific type.
// - Only accepts the send operation "ch <- value".

var _ chan<- int

// Declaring a Receive-Only Channel
// - A receive-only channel can only be used to receive values of a specific type.
// - Only accepts the receive operation "<-ch".

var _ <-chan int

//==================================================================================================
// Declaring Goroutines
//==================================================================================================

// Declaring a Goroutine
// - A goroutine is a lightweight thread managed by the Go runtime.
// - It is used to run functions concurrently.
// - The "go" keyword followed by a function call starts a new goroutine.

func StartGoroutines() {

	// Start a Goroutine
	// - This will run the print function concurrently.
	// - Output: Hello

	go fmt.Println("Hello")

	// Start a Goroutine with Function
	// - We can also start a goroutine with a function.
	// - Note that the function must be called with parentheses.
	// - Output: Hello

	go func() {
		fmt.Println("Hello")
	}()
}

//==================================================================================================
// Communicating With Channels
//==================================================================================================

// Communicating with Channels
// - Channels are used to communicate between goroutines.
// - They can be thought of as pipes that connect goroutines.

func CommunicatingWithChannels() {

	// Declaring a Channel
	// - We will declare a channel to be used for communication.
	// - We can use the "make" function to create a channel.

	ch := make(chan int)

	// Start a Goroutine with a Channel
	// - Note that the goroutine will send a value to the channel.

	go func() {
		ch <- 42
	}()

	// Receiving from a Channel
	// - We can receive a value from the channel using the "<-ch" syntax.
	// - Note: It will lock the current thread until a value is received from the channel.
	// - Output: 42

	x := <-ch
	fmt.Println(x)
}

//==================================================================================================
// Using Multiple Channels
//==================================================================================================

// Using Multiple Channels (Select)
// - The select statement is used to wait on multiple channel operations.
// - It can also be used to avoid thread blocking when waiting for a channel to be ready.
// - It is similar to a switch statement, but for channels.

func UsingMultipleChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Start Goroutines with Channels
	// - Note that the goroutines will send values to the channels.

	go func() {
		ch1 <- 4
	}()
	go func() {
		ch2 <- 8
	}()

	// Using Select Statement
	// - We can use the select statement to wait on multiple channel operations.
	// - The value received depends on which channel is ready first.
	// - Output: Received from ch1: 4

	select {
	case x := <-ch1:
		fmt.Println("Received from ch1:", x)
	case y := <-ch2:
		fmt.Println("Received from ch2:", y)
	}
}

//==================================================================================================
// Skip Channel Waiting
//==================================================================================================

// Skip Channel Waiting
// - We can use the "select" statement to skip waiting for a channel to be ready.

func SkipChannelWaiting() {
	ch := make(chan int)

	// Start a Goroutine with a Channel
	// - We will sleep for 1 second before sending a value to the channel.

	go func() {
		time.Sleep(time.Second)
		ch <- 42
	}()

	// Using Select Statement
	// - We can use the select statement to skip waiting for a channel to be ready.
	// - If the channel is not ready, the default case will be executed.
	// - Output: No value received from ch | Received from ch: 42

	for {
		select {
		case x := <-ch:
			fmt.Println("Received from ch:", x)
			return
		default:
			fmt.Println("No value received from ch")
			time.Sleep(time.Second)
		}
	}
}

//==================================================================================================
// Unbuffered Channels
//==================================================================================================

// Unbuffered Channels
// - Unbuffered channels are used for synchronous communication between goroutines.
// - They block the sending goroutine until the receiving goroutine is ready to receive the value.

func UnbufferedChannels() {
	ch := make(chan int)

	// Start a Goroutine with a Channel
	// - Note that two values are sent to the channel.

	go func() {
		fmt.Println("Send: 1")
		ch <- 1
		fmt.Println("Send: 2")
		ch <- 2
	}()

	// Receiving from a Channel
	// - Note that the main goroutine will block until a value is received from the channel.
	// - The goroutine will also block until the main goroutine is ready to receive the value.
	// - This is the key feature of unbuffered channels.
	// - Output: Send: 1 | Received: 1 | Send: 2 | Received: 2

	time.Sleep(time.Second)
	x := <-ch
	fmt.Println("Received:", x)
	time.Sleep(time.Second)
	x = <-ch
	fmt.Println("Received:", x)
}

//==================================================================================================
// Buffered Channels
//==================================================================================================

// Buffered Channels
// - Buffered channels are used for asynchronous communication between goroutines.
// - They allow sending and receiving values without blocking the sending goroutine until the buffer
//   is full.

func BufferedChannels() {

	// Declaring a Buffered Channel
	// - To create a buffered channel, we can use the "make" function with a buffer size.

	ch := make(chan int, 2)

	// Start a Goroutine with a Buffered Channel
	// - Note that two values are sent to the channel.

	go func() {
		fmt.Println("Sent: 1")
		ch <- 1
		fmt.Println("Sent: 2")
		ch <- 2
	}()

	// Receiving from a Buffered Channel
	// - Now, since we have a buffered channel, the main goroutine will not block until the buffer
	//   is full.
	// - The goroutine will also not block until the buffer is full.
	// - Output: Send: 1 | Send: 2 | Received: 1 | Received: 2

	time.Sleep(time.Second)
	x := <-ch
	fmt.Println("Received:", x)
	time.Sleep(time.Second)
	x = <-ch
	fmt.Println("Received:", x)
}
