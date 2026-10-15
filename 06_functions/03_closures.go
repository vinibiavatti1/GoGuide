// Closures
// - Closures are functions that capture and reference variables from outside their immediate body.
// - They "close over" these variables, allowing them to read or modify external state even after
//   the outer function has finished executing.
// - Captured variables become "private" state, accessible only through the closure itself.

package functions

import "fmt"

// Closure Function
// - The closure function below declares a "private" variable, keeping state across multiple
//   invocations.

func Counter() func() int {
	x := -1
	return func() int {
		x++
		return x
	}
}

// Usage
// - In this example, we use the closure above to create a counter.
// - Note that the closure keeps the state of the counter across multiple invocations.
// - This happens because the variable `x` escapes to the heap and persists in memory as long as the
//   closure references it.
// - Output: 0 1 2

func UsingClosures() {
	next := Counter()
	fmt.Println(next(), next(), next())
}
