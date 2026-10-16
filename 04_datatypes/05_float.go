// Float
// - The float type in Go is used to represent floating-point numbers.
// - It is used to store decimal values with variable precision.
// - Go provides two primary types for floating-point numbers:
//   - float32: a 32-bit floating-point number
//   - float64: a 64-bit floating-point number (default type)

package datatypes

import "fmt"

//==================================================================================================
// Declaring Floats
//==================================================================================================

func DeclaringFloats() {

	// Declaring a Float
	// - We can declare a float using a floating-point number.
	// - The default type is float64.
	// - Output: x: 3.14

	x := 3.14
	fmt.Println("x:", x)

	// Creating Float with Scientific Notation
	// - The scientific notation ("e" or "E") can be used to represent floats.
	// - Note that the result is a float64.
	// - Output: x: 12300

	x = 1.23e4
	fmt.Println("x:", x)

	// Creating a 32-bit Float
	// - We can explicitly declare a float32 type.
	// - Output: y: 3.14

	var y float32 = 3.14
	fmt.Println("y:", y)

	// Float Precision
	// - Floating-point numbers are not always perfectly accurate due to the limitations of binary
	//   floating-point representation.
	// - Note that in the example below, the result is not exactly 0.3.
	// - Output: x: 0.30000000000000004

	x = 0.1 + 0.2
	fmt.Println("x:", x)
}

//==================================================================================================
// Float Operations
//==================================================================================================

func FloatOperations() {

	// Addition
	// - We can use the "+" operator to add two floats.
	// - Output: x: 5.5

	x := 3.0 + 2.5
	fmt.Println("x:", x)

	// Subtraction
	// - We can use the "-" operator to subtract two floats.
	// - Output: x: 0.5

	x = 3.0 - 2.5
	fmt.Println("x:", x)

	// Multiplication
	// - We can use the "*" operator to multiply two floats.
	// - Output: x: 7.5

	x = 3.0 * 2.5
	fmt.Println("x:", x)

	// Division
	// - We can use the "/" operator to divide two floats.
	// - Note: Modulo operation is not supported for floating-point numbers.
	// - Output: x: 2.75

	x = 5.5 / 2.0
	fmt.Println("x:", x)
}

//==================================================================================================
// Float Conversion
//==================================================================================================

func FloatConversion() {

	// Widening conversion (expanding the size)
	// - This example expands the size from 32-bit to 64-bit.
	// - Output: float32: 1.5 float64: 1.5

	var x1 float32 = 1.5
	var y1 float64 = float64(x1)
	fmt.Println("float32:", x1, "float64:", y1)

	// Narrowing conversion (truncating excess bits)
	// - This example reduces the size from 64-bit to 32-bit.
	// - Output: float64: 3.14 float32: 3.14

	var x2 float64 = 3.14
	var y2 float32 = float32(x2)
	fmt.Println("float64:", x2, "float32:", y2)

	// Converting float to int
	// - Floats can be converted to integers by truncating the decimal part.
	// - Output: float64: 3.99 int: 3

	var x3 float64 = 3.99
	var y3 int = int(x3)
	fmt.Println("float64:", x3, "int:", y3)

	// Converting int to float
	// - Integers can be converted to floats without losing precision.
	// - Output: int: 5 float64: 5

	var x4 int = 5
	var y4 float64 = float64(x4)
	fmt.Println("int:", x4, "float64:", y4)

	// String
	// - Convert to string using the "fmt" package.
	// - Output: 3.14
	s1 := fmt.Sprint(3.14)
	fmt.Println(s1)

	// Formatted String
	// - We can convert a float specifying precision using format verbs like "%.2f".
	// - Output: 3.14
	s2 := fmt.Sprintf("%.2f", 3.1415)
	fmt.Println(s2)
}
