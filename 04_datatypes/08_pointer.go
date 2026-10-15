// Pointer
// - Pointers are a way to reference a variable's memory address in Go.
// - They allow you to modify the value of a variable without passing it by value.
// - Pointers are declared using the `*` operator, and you can obtain a pointer to a variable using
//   the `&` operator.
// - Syntax:
//   - *x - dereference operator, used to access the value at the address stored in x
//   - &x - address operator, used to get the memory address of x

package datatypes

import "fmt"

//==================================================================================================
// Using Pointers
//==================================================================================================

func UsingPointers() {
	x := 5

	// Declaring a Pointer
	// - We will declare a pointer to hold the address of x.
	// - We can use the "&" operator to get the address of x.
	// - Output: p: 0xc00000a0b8

	p := &x
	fmt.Println("p:", p)

	// Dereferencing a Pointer
	// - We can use the "*" operator to dereference the pointer and get the value at the address.
	// - Output: y: 5

	y := *p
	fmt.Println("y:", y)

	// Changing the Value of a Pointer
	// - Functions can receive pointers instead of values, to manipulate the original variable.
	// - Output: x: 10

	changeValue(p)
	fmt.Println("x:", *p)

	// Pointers Can Have Nil Values
	// - A pointer can be nil, meaning it doesn't point to any address.
	// - Output: p: <nil>

	p = nil
	fmt.Println("p:", p)
}

//==================================================================================================
// Pointer Reference
//==================================================================================================

// Pointers
// - Pointers hold the memory address of a value.
// - Passing a pointer allows functions to modify the original data without copying it.
// - Changes made through the pointer will affect the original variable.
// - In the example below, the "x" variable will be modified outside the function.

func changeValue(x *int) {
	*x = 10
}
