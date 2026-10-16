// Pointer
// - Pointers are a way to reference a variable's memory address in Go.
// - They allow you to modify the value of a variable without passing it by value.
// - Pointers are declared using the `*` operator, and you can obtain a pointer to a variable using
//   the `&` operator.
// - Syntax:
//   - *<type> pointer type declaration, used to define that a variable or field holds a reference
//   - *<var>  dereference operator, used to access the value at the address stored in x
//   - &<var>  address operator, used to get the memory address of x«

package datatypes

import "fmt"

//==================================================================================================
// Pointers
//==================================================================================================

func Pointers() {

	// Declaring a Pointer Type (* in type)
	// - To declare a pointer type, the "*" is used on the data type (e.g., "*int").
	// - Note: Pointers can store nil values when they are not initialized to any memory address.
	// - Output: ptr: <nil>

	var ptr *int = nil
	fmt.Println("ptr:", ptr)

	// Address Operator (& in variable)
	// - To get the memory address of a variable, the "&" operator is used before the variable name.
	// - Output: ptr: 0xc00000a0b8 (hexadecimal memory address)

	x := 5
	ptr = &x
	fmt.Println("ptr:", ptr)

	// Dereferencing a Pointer (* in variable)
	// - To access or retrieve the actual value stored at the memory address, the "*" operator is
	//   used on the pointer variable.
	// - Output: y: 5

	y := *ptr
	fmt.Println("y:", y)
}

//==================================================================================================
// Pointer Arguments
//==================================================================================================

// Pointers Arguments
// - Passing a pointer arguments allow functions to modify the original data without copying it.
// - Changes made through the pointer will affect the original variable.

func SetValue(x *int) {
	*x = 10
}

// Example
// - In the example below, the "x" variable will be modified outside the function.
// - Output: x: 10

func ProcessValue() {
	x := 5
	SetValue(&x)
	fmt.Println("x:", x)
}
