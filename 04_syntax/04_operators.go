// Operators
// Go provides a rich set of operators to work with different types of data:
// - Arithmetic Operators: +, -, *, /, %
// - Assignment Operators: =, +=, -=, *=, /=, %=
// - Comparison Operators: ==, !=, >, <, >=, <=
// - Logical Operators:    &&, ||, !
// - Address Operators:    & *
// - Bitwise Operators:    &, |, ^, <<, >>

package syntax

import "fmt"

// Arithmetic Operators
// These operators perform arithmetic operations such as addition, subtraction,
// multiplication, and division.
func ArithmeticOperators() {
	// Creating variables
	// We will use these variables to demonstrate arithmetic operators.
	x, y := 10, 5

	// Arithmetic Operators
	_ = x + y // 15 Addition
	_ = x - y // 5  Subtraction
	_ = x * y // 50 Multiplication
	_ = x / y // 2  Division
	_ = x % y // 0  Modulo (Remainder)
}

// Assignment Operators
// These operators are used to assign values to variables.
func AssignmentOperators() {
	// Creating variables
	// We will use these variables to demonstrate assignment operators.
	x := 10

	// Assignment Operators
	x = 20 // 20 Assignment
	x += 5 // 25 Add and Assign
	x -= 3 // 22 Subtract and Assign
	x *= 2 // 44 Multiply and Assign
	x /= 4 // 11 Divide and assign
	x %= 3 // 2  Modulo and assign
}

// Comparison Operators
// These operators compare values and return a boolean result (true or false).
func ComparisonOperators() {
	// Creating variables
	// We will use these variables to demonstrate comparison operators.
	x, y := 10, 5

	// Comparison Operators
	_ = x == y // false Equal
	_ = x != y // true  Not Equal
	_ = x > y  // true  Greater Than
	_ = x < y  // false Less Than
	_ = x >= y // true  Greater Than or Equal To
	_ = x <= y // false Less Than or Equal To
}

// Logical Operators
// These operators perform logical operations, typically used with boolean values.
func LogicalOperators() {
	// Creating boolean variables
	// We will use these variables to demonstrate logical operators.
	x, y := true, false

	// Logical Operators
	_ = x && y // false AND
	_ = x || y // true  OR
	_ = !x     // false NOT
}

// Address Operators
// These operators are used to get the memory address of a variable
// (using "&") and to dereference pointers (using "*").
func AddressOperators() {
	// Creating a variable
	// We will use this variable to demonstrate address operators.
	x := 10

	// Creating a pointer
	// Getting the memory address of "a" (address operator "&")
	ptr := &x
	fmt.Println("Address:", ptr) // Output: Address: 0xc000067f30

	// Dereferencing a Pointer
	// Getting the value stored at the memory address (dereference operator "*")
	fmt.Println("Value:", *ptr) // Output: Value: 10
}

// Bitwise Operators
// These operators perform operations at the bit level on integer values.
func BitwiseOperators() {
	// Creating variables
	// We will use these variables to demonstrate bitwise operators.
	x := 60 // 0011 1100
	y := 13 // 0000 1101

	// Bitwise Operators
	_ = ^x     // 1100 0011 NOT
	_ = x & y  // 0000 1100 AND
	_ = x | y  // 0011 1101 OR
	_ = x ^ y  // 0011 0001 XOR
	_ = x &^ y // 0011 0000 AND NOT (Bit Clear)
	_ = x << 2 // 1111 0000 Left Shift
	_ = x >> 2 // 0000 1111 Right Shift
}
