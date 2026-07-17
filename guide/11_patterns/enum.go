// Enum
// Go does not have a built-in enum type, but we can use a custom type with constants to achieve similar functionality.
// We can define a custom type and use iota to create a set of related constants.
// This is a common pattern in Go to create enumerated types.

package patterns

import "fmt"

// Defining a Custom Type
// We will define a custom type to represent out enumerated type.
type Color int

// Defining Constants
// We will use iota to create a set of data values.
const (
	Red   Color = iota // 0
	Green              // 1
	Blue               // 2
)

// Defining an Array with All Enum Values
// We can define an array to hold all the enum values.
var Colors = [3]Color{Red, Green, Blue}

// String Map
// We will define a map to convert the enum values to string representations.
var colorToString = map[Color]string{
	Red:   "Red",
	Green: "Green",
	Blue:  "Blue",
}

// String Method
// We will implement a String method to convert the enum value to a string representation.
func (c Color) String() string {
	if str, ok := colorToString[c]; ok {
		return str
	}
	return "Unknown"
}

// Using Enum
// This function demonstrates how to use the enum type and its methods.
func UsingEnum() {

	// Get Enum Value
	// We can get the enum value by accessing the enum constant directly.
	x := Red
	fmt.Println("x:", x) // Output: x: 0

	// Get Enum String Value
	// We can get the string representation of the enum value by calling the String method.
	fmt.Println("x:", x.String()) // Output: x: Red

	// Iterate Over Enum Values
	// We can iterate over the enum values using a for loop.
	for _, c := range Colors {
		fmt.Println(c, c.String()) // Output: 0 Red, 1 Green, 2 Blue
	}
}
