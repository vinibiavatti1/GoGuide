// Integer
// - An integer (int) is a fundamental numeric datatype in Go, used to represent integer numbers.
// - Signed integers can represent positive and negative numbers.
// - Unsigned integers can only represent positive numbers.
// - The byte and rune types are aliases for integer types in Go.

package datatypes

import "fmt"

//==================================================================================================
// Declaring Integers
//==================================================================================================

func DeclaringIntegers() {

	// Declaring an Integer
	// - The "int" type is used to represent integers.
	// - The size depends on the platform (32 or 64 bits).
	// - Output: x: 1

	x := 1
	fmt.Println("x:", x)

	// Declaring an Integer Using Binary Notation
	// - We can use the "0b<n>" notation to define integers in binary.
	// - Output: x: 10

	x = 0b1010
	fmt.Println("x:", x)

	// Declaring an Integer Using Hexadecimal Notation
	// - We can use the "0x<n>" notation to define integers in hexdecimal.
	// - Output: x: 4095

	x = 0xFFF
	fmt.Println("x:", x)

	// Declaring an Integer Using Octal Notation
	// - We can use the "0o<n>" notation to define integers in octal.
	// - Output: x: 511

	x = 0o777
	fmt.Println("x:", x)

	// Declaring an Integer Using a Rune
	// - Since a rune is an alias for int32, we can use it to represent integers.
	// - A rune is a Unicode code point.
	// - Output: x: 65

	x = 'A'
	fmt.Println("x:", x)

	// Declaring an Integer Using Underscores
	// - We can use the underscore "_" to separate digits and improve readability.
	// - Output: x: 100000000

	x = 100_000_000
	fmt.Println("x:", x)

	// Declaring an Computed Integer
	// - We can create integers from expressions which will be computed to generate the result.
	// - Output: x: 14

	x = 2 + 3*4
	fmt.Println("x:", x)

	// Declaring other Integer Types
	// - Other integer types can be used to represent different sizes.
	// - The unsigned integer types can only represent positive numbers.

	var _ int8 = 1    // 8-bit: -128 to 127
	var _ int16 = 1   // 16-bit: -32768 to 32767
	var _ int32 = 1   // 32-bit: -2147483648 to 2147483647
	var _ int64 = 1   // 64-bit: -9223372036854775808 to 9223372036854775807
	var _ uint = 1    // 32/64-bit Unsigned integer (default)
	var _ uint8 = 1   // 8-bit: 0 to 255
	var _ uint16 = 1  // 16-bit: 0 to 65535
	var _ uint32 = 1  // 32-bit: 0 to 4294967295
	var _ uint64 = 1  // 64-bit: 0 to 18446744073709551615
	var _ uintptr = 1 // Memory Address Type (32 or 64 bits)
	var _ byte = 1    // Alias for uint8
	var _ rune = 1    // Alias for int32 (Unicode code point)
}

//==================================================================================================
// Integer Operations
//==================================================================================================

func IntegerOperations() {

	// Addition
	// - The "+" operator is used to add two integers.
	// - Output: x: 10

	x := 7 + 3
	fmt.Println("x:", x)

	// Subtraction
	// The "-" operator is used to subtract two integers.
	// Output: x: 4

	x = 7 - 3
	fmt.Println("x:", x)

	// Multiplication
	// The "*" operator is used to multiply two integers.
	// Output: x: 21

	x = 7 * 3
	fmt.Println("x:", x)

	// Division
	// The "/" operator is used to divide two integers.
	// Output: x: 3

	x = 9 / 3
	fmt.Println("x:", x)

	// Modulo Operation (Remainder)
	// The "%" operator is used to get the remainder of a division.
	// Output: x: 1

	x = 3 % 2
	fmt.Println("x:", x)
}

//==================================================================================================
// Integer Conversions
//==================================================================================================

func IntegerConversions() {

	// Widening conversion (expanding the size)
	// - This example expands the size from 8-bit to 16-bit.
	// - Output: int8: 127 int16: 127

	var x1 int8 = 127
	var y1 int16 = int16(x1)
	fmt.Println("int8:", x1, "int16:", y1)

	// Narrowing conversion (truncating excess bits)
	// - This example reduces the size from 16-bit to 8-bit.
	// - Output: int16: 32767 int8: -1

	var x2 int16 = 32767
	var y2 int8 = int8(x2)
	fmt.Println("int16:", x2, "int8:", y2)

	// Signed to unsigned conversion
	// - This example changes an signed integer to unsigned.
	// - Output: int8: -1 uint8: 255

	var x3 int8 = -1
	var y3 uint8 = uint8(x3)
	fmt.Println("int8:", x3, "uint8:", y3)

	// Unsigned to signed conversion
	// - This example changes an unsigned integer to signed.
	// - Output: uint8: 255 int8: -1

	var x4 uint8 = 255
	var y4 int8 = int8(x4)
	fmt.Println("uint8:", x4, "int8:", y4)

	// String
	// - Convert to string using the "fmt" package.
	// - Output: 128
	str := fmt.Sprint(128)
	fmt.Println(str)
}
