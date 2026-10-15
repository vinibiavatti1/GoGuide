// Defer
// - A defer statement defers the execution of a function until the surrounding function returns.
// - The deferred call's arguments are evaluated immediately, but the function call is not executed
//   until the surrounding function returns.
// - The execution order of deferred functions is LIFO (Last In First Out).
// - Deferred functions avoid needing to write closing syntax, such as closing a file or
//   unlocking a mutex for each code case.

package functions

import "fmt"

// Single Defer
// - A defer statement defers the execution of a function until the surrounding function returns.
// - This exmaple shows a single defer statement.
// - The End message is printed before the Deferred message.
// - Output: End | Defer

func SingleDefer() {
	defer fmt.Println("Defer")
	fmt.Println("End")
}

// Multiple Defers
// - The deferred functions are executed in reverse order.
// - The execution order of deferred functions is LIFO (Last In First Out).
// - Output: End | Defer 2 | Defer 1

func MultipleDefers() {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	fmt.Println("End")
}

// Defer Arguments
// - The deferred call's arguments are evaluated immediately, but the function call is not executed
//   until the surrounding function returns.
// - Notice that the deferred function prints "Deferred" instead of "End".
// - Output: Deferred

func DeferArguments() {
	x := "Deferred"
	defer fmt.Println(x)
	x = "End"
}

// Defer Loop
// - The deferred function is executed after the loop ends.
// - As we can see, the values of i are printed in reverse order.
// - Output: 2 | 1 | 0

func DeferLoop() {
	for i := range 3 {
		defer fmt.Println(i)
	}
}

// Defer with Function
// - We can defer a function execution.
// - The deferred function is executed after the surrounding function returns.
// - Notice that the function must be Immediately Invoked.
// - Output: End | Function

func DeferFunction() {
	defer func() { fmt.Println("Function") }()
	fmt.Println("End")
}

// Defer Named Return
// - A deferred function can modify the return value of a named return.
// - The deferred function modifies the return value of x.
// - Output: 15

func DeferNamedReturn() (x int) {
	x = 5
	defer func() { x += 10 }()
	return
}
