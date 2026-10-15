// Math
// The math package provides basic constants and mathematical functions.
// All functions in this package work with float64 types, requiring explicit type casting for integers.

package library

import (
	"fmt"
	"math"
)

// Basic Constants
// The following constants represent common mathematical properties.
func BasicConstants() {
	// Pi
	// Pi is the mathematical constant ratio of a circle's circumference to its diameter.
	// Output: 3.141592653589793
	fmt.Println(math.Pi)

	// E
	// E is Euler's number, the base of natural logarithms.
	// Output: 2.718281828459045
	fmt.Println(math.E)
}

// Basic Functions
// The following functions perform fundamental mathematical operations.
func BasicFunctions() {
	// Abs
	// Abs returns the absolute value of x (turns negative numbers into positive).
	// Output: 42
	fmt.Println(math.Abs(-42.0))

	// Max
	// Max returns the larger of x or y.
	// Output: 10
	fmt.Println(math.Max(5.5, 10.0))

	// Min
	// Min returns the smaller of x or y.
	// Output: 5.5
	fmt.Println(math.Min(5.5, 10.0))

	// Pow
	// Pow returns x**y (x raised to the power of y).
	// Output: 8 (2 cubed)
	fmt.Println(math.Pow(2, 3))

	// Sqrt
	// Sqrt returns the square root of x.
	// Output: 3
	fmt.Println(math.Sqrt(9))
}

// Rounding Functions
// The following functions provide different ways to round floating-point numbers.
func RoundingFunctions() {
	// Ceil
	// Ceil returns the least integer value greater than or equal to x (rounds UP).
	// Output: 4
	fmt.Println(math.Ceil(3.14))

	// Floor
	// Floor returns the greatest integer value less than or equal to x (rounds DOWN).
	// Output: 3
	fmt.Println(math.Floor(3.99))

	// Round
	// Round returns the nearest integer, rounding half away from zero.
	// Output: 4 3
	fmt.Println(math.Round(3.5), math.Round(3.4))
}

// Trigonometric Functions
// The following functions perform standard trigonometric calculations.
// Note: Go expects angles to be measured in radians, not degrees.
func TrigonometricFunctions() {
	// Sin
	// Sin returns the sine of the radian argument x.
	// Output: 0 (or a value extremely close to 0 due to floating-point precision)
	fmt.Println(math.Sin(math.Pi))

	// Cos
	// Cos returns the cosine of the radian argument x.
	// Output: -1
	fmt.Println(math.Cos(math.Pi))

	// Tan
	// Tan returns the tangent of the radian argument x.
	// Output: 0
	fmt.Println(math.Tan(0))

	// Converting Degrees to Radians
	// Since Go uses radians, to find the sine of 90 degrees you must convert it first.
	// Output: 1
	degrees := 90.0
	radians := degrees * (math.Pi / 180)
	fmt.Println(math.Sin(radians))
}
