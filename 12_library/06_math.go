// Math
// - The math package provides basic constants and mathematical functions.
// - All functions in this package work with float64 types, requiring explicit type casting for
//   integers.

package library

import (
	"fmt"
	"math"
)

func MathOperations() {

	// Constants

	fmt.Println(math.E)      // 2.718281828459045
	fmt.Println(math.Pi)     // 3.141592653589793
	fmt.Println(math.Phi)    // 1.618033988749895 (Golden ratio)
	fmt.Println(math.Log10E) // 0.4342944819032518
	fmt.Println(math.Log2E)  // 1.4426950408889634
	fmt.Println(math.Sqrt2)  // 1.4142135623730951
	fmt.Println(math.SqrtE)  // 1.6487212707001282
	fmt.Println(math.Pi)     // 3.141592653589793

	// Rounding & Truncation

	fmt.Println(math.Ceil(4.1))        // 5
	fmt.Println(math.Floor(4.9))       // 4
	fmt.Println(math.Round(4.5))       // 5
	fmt.Println(math.RoundToEven(4.5)) // 4 (Round ties to even)
	fmt.Println(math.Trunc(4.9))       // 4

	// Basic Arithmetic & Utilities

	fmt.Println(math.Abs(-10.0))          // 10
	fmt.Println(math.Min(5.0, 1.0))       // 1
	fmt.Println(math.Max(5.0, 10.0))      // 10
	fmt.Println(math.Dim(10.0, 4.0))      // 6 (Max(x-y, 0))
	fmt.Println(math.Mod(5.0, 2.0))       // 1 (Floating-point remainder)
	fmt.Println(math.Remainder(5.0, 3.0)) // -1 (IEEE 754 remainder)
	fmt.Println(math.Copysign(5.0, -1.0)) // -5 (Returns magnitude of x with sign of y)

	// Exponentials, Powers & Roots

	fmt.Println(math.Pow(2.0, 3.0))   // 8
	fmt.Println(math.Pow10(2))        // 100 (10**n)
	fmt.Println(math.Sqrt(16.0))      // 4
	fmt.Println(math.Cbrt(27.0))      // 3
	fmt.Println(math.Exp(1.0))        // 2.718281828459045 (e^1)
	fmt.Println(math.Expm1(1.0))      // 1.718281828459045 (e^1 - 1)
	fmt.Println(math.Hypot(3.0, 4.0)) // 5 (sqrt(3^2 + 4^2))

	// Logarithms

	fmt.Println(math.Log(math.E))  // 1 (natural log)
	fmt.Println(math.Log10(100.0)) // 2 (base 10)
	fmt.Println(math.Log2(8.0))    // 3 (base 2)
	fmt.Println(math.Log1p(0.0))   // 0 (ln(1 + 0))

	// Trigonometry

	fmt.Println(math.Sin(math.Pi / 2.0)) // 1 (90 deg)
	fmt.Println(math.Cos(math.Pi))       // -1 (180 deg)
	fmt.Println(math.Tan(math.Pi / 4.0)) // 1 (45 deg)
	fmt.Println(math.Asin(1.0))          // 1.5707963267948966 (Pi / 2 radians)
	fmt.Println(math.Acos(-1.0))         // 3.141592653589793 (Pi radians)
	fmt.Println(math.Atan(1.0))          // 0.7853981633974483 (Pi / 4 radians)
	fmt.Println(math.Atan2(5.0, 5.0))    // 0.7853981633974483 (angle for point y=5, x=5)
	fmt.Println(math.Sincos(math.Pi))    // -1 1.2246467991473532e-16 (Returns Sin and Cos)

	// Hyperbolic Trigonometry

	fmt.Println(math.Sinh(0.0))  // 0
	fmt.Println(math.Cosh(0.0))  // 1
	fmt.Println(math.Tanh(0.0))  // 0
	fmt.Println(math.Asinh(0.0)) // 0
	fmt.Println(math.Acosh(1.0)) // 0
	fmt.Println(math.Atanh(0.0)) // 0

	// Float Checks & Analysis

	fmt.Println(math.IsNaN(math.NaN()))     // true
	fmt.Println(math.IsInf(math.Inf(1), 1)) // true (Positive infinity check)
	fmt.Println(math.Signbit(-5.0))         // true (Reports if negative or negative zero)
}
