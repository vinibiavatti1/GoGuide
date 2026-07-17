// Any
// The "any" data type is used to store values of any type.
// It is equivalent to the "interface{}" type in Go.
// The "any" type is useful when you need to store values of different types in a single variable.

package datatypes

import "fmt"

// Declaring Any Type
// The "any" type can store values of any type.
func DeclaringAnyType() {
	// Declaring a Variable of Type "any"
	// The "any" type can store values of any type.
	// Output: 1
	var x any = 1
	fmt.Println(x)

	// Assigning a Different Type
	// We can assign values of different types to the same variable.
	x = "Hello, World!"
	x = true
	x = 3.14

	// Type Assertion
	// We can use type assertion to extract the underlying value from an "any" type.
	// The syntax is: value, ok := x.(T)
	// If the type assertion fails, "ok" will be false.
	// Output: 3.14 true
	y, ok := x.(float64)
	fmt.Println(y, ok)

	// Type Switch with "any"
	// A type switch is used to compare the type of an interface value.
	// It is similar to a type assertion, but with a switch statement.
	// (Matched) Output: x is a float64
	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case float64:
		fmt.Println("x is a float64")
	}
}
