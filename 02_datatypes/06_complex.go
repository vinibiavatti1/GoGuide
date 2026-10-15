// Complex
// - A complex number has two parts: a real part and an imaginary part.
// - The real and imaginary parts are floating-point numbers.
// - Go provides two types for complex numbers:
//   - complex64: a complex number where both real and imaginary parts are of type float32.
//   - complex128: (default type) a complex number where both real and imaginary parts are of type
//     float64.

package datatypes

import "fmt"

//==================================================================================================
// Declaring Complex Numbers
//==================================================================================================

func DeclaringComplexNumbers() {

	// Declaring a Complex Number
	// - The "i" notation is used to define the imaginary part of a complex number.
	// - The complex128 is the default type for complex numbers.
	// - Output: x: (1+2i)
	x := 1 + 2i
	fmt.Println("x:", x)

	// Declaring a Complex Number (Using "complex()")
	// - The "complex()" function takes two float64 arguments: real and imaginary parts.
	// - Output: x: (1+2i)
	x = complex(1, 2)
	fmt.Println("x:", x)

	// Declaring a Complex Number of 64-bit Type
	// - We can specify the type of a complex number by using the complex64 type.
	// - Output: y: (1+2i)
	var y complex64 = complex(1, 2)
	fmt.Println("y:", y)
}

//==================================================================================================
// Manipulating Complex Numbers
//==================================================================================================

func ManipulatingComplexNumbers() {
	x := 1 + 2i

	// Acessing Real Part
	// - We can use the "real()" builtin function to access the real part of a complex number.
	// - Output: Real Part: 1
	r := real(x)
	fmt.Println("Real Part:", r)

	// Acessing Imaginary Part
	// - We can use the "imag()" builtin function to access the imaginary part of a complex number.
	// - Output: Imaginary Part: 2
	i := imag(x)
	fmt.Println("Imaginary Part:", i)
}

//==================================================================================================
// Complex Number Operations
//==================================================================================================

func ComplexNumberOperations() {
	x := 1 + 2i
	y := 3 + 4i

	// Addition
	// - The "+" operator is used to add two complex numbers.
	// - Output: z: (4+6i)
	z := x + y
	fmt.Println("z:", z)

	// Subtraction
	// - The "-" operator is used to subtract two complex numbers.
	// - Output: z: (-2-2i)
	z = x - y
	fmt.Println("z:", z)

	// Multiplication
	// - The "*" operator is used to multiply two complex numbers.
	// - Output: z: (-5+10i)
	z = x * y
	fmt.Println("z:", z)

	// Division
	// - The "/" operator is used to divide two complex numbers.
	// - Note: Modulo operation is not supported for complex numbers.
	// - Output: z: (0.44+0.08i)
	z = x / y
	fmt.Println("z:", z)
}

//==================================================================================================
// Complex Numbers Conversion
//==================================================================================================

func ComplexNumbersConversion() {

	// Widening conversion (expanding the size)
	// - This example expands the size from 64-bit to 128-bit.
	// - Output: complex64: (1+2i) complex128: (1+2i)
	var x1 complex64 = complex(1, 2)
	var y1 complex128 = complex128(x1)
	fmt.Println("complex64:", x1, "complex128:", y1)

	// Narrowing conversion (truncating excess bits)
	// - This example reduces the size from 128-bit to 64-bit.
	// - Output: complex128: (3+4i) complex64: (3+4i)
	var x2 complex128 = complex(1, 2)
	var y2 complex64 = complex64(x2)
	fmt.Println("complex128:", x2, "complex64:", y2)
}
