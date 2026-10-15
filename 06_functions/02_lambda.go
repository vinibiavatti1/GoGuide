// Lambda
// - Lambda functions are anonymous functions that can be assigned to variables.
// - Syntax: func(<parameters> <type>...) <return> {...}

package functions

import "fmt"

func LambdaFunctions() {

	// Lambda Function
	// - An example of a lambda function without parameters and return values.
	// - Output: Hello, World!

	fn1 := func() {
		fmt.Println("Hello, World!")
	}
	fn1()

	// Lambda Function with Parameters and Return Value
	// - An example of a lambda function with parameters and a return values.
	// - Output: x: 4

	fn2 := func(x int) int {
		return x * 2
	}
	x := fn2(2)
	fmt.Println("x:", x)

	// Immediately Invoked Function Expression (IIFE)
	// - An example of an immediately invoked lambda function.
	// - This is useful for deferred execution.
	// - Output: 6

	x = func(y int) int {
		return y * 2
	}(3)
	fmt.Println("x:", x)
}
