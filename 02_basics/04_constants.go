// Constants
// - In Go, constants are values that cannot be changed once set. Constants can be of any type,
//   including numeric, string, or boolean types.
// - Constants must be declared before they can be used.
// - Note: The value of a constant must be determinable at compile time.

package basics

import "fmt"

// Declaring Global Constants
// - We can declare global constants outside of functions.
// - Global constants can be accessed by any function inside the package.
const Const1 int = 1        // Type explicitly
const Const2 = 2            // Type Inference
const Const3, Const4 = 3, 4 // Inline
const _ int = 1             // Unused

// Declaring Multiple Global Constants
// - We can declare multiple global constants in multiple lines using the const keyword.
const (
	Const5         int = 1    // Type explicitly
	Const6             = 2    // Type Inference
	Const7, Const8     = 3, 4 // Inline
	_                  = 1    // Unused
)

// Declaring Multiple Global Constants With Literals
// - We can repeat the last literal value in a constant block.
// - This is useful when we want to declare a series of constants with the same value.
// - Note: This is not possible with variables.
const (
	ConstA = 1 // Value: 1
	ConstB     // Value: 1 (repeated)
	ConstC = 2 // Value: 2
	ConstD     // Value: 2 (repeated)
)

//==================================================================================================
// Declaring Constants
//==================================================================================================

func DeclaringConstants() {

	// Declaring a Constant
	// - In this example, we declare a constant with the type int explicitly.
	// - Output: x1: 10
	const x1 int = 10
	fmt.Println("x1:", x1)

	// Declaring a constant with Type Inference
	// - The compiler can infer the type of a constant based on the value assigned to it.
	// - Output: x2: 10
	const x2 = 10
	fmt.Println("x2:", x2)

	// Declaring Unused constants
	// - We can use the blank identifier (_) to declare unused constants.
	const _ int = 1
	const _ = 1

	// Declaring Inline constants
	// - We can declare multiple constants in a single line.
	// - Note that the last type declared is used for all constants.
	// - Output: Inline: 1 2 3
	const a1, b1, c1 int = 1, 2, 3
	fmt.Println("Inline:", a1, b1, c1)

	// Declaring Inline constants With Type Inference
	// - We can declare multiple constants in a single line using type inference.
	// - Output: Inline: 1 true 3.5
	const a2, b2, c2 = 1, true, 3.5
	fmt.Println("Inline:", a2, b2, c2)

	// Declaring Multiple constants
	// - We can declare multiple constants in multiple lines using the const keyword.
	// - Output: Consts: 1 2 3 4
	const (
		const1         int = 1    // Type explicitly
		const2             = 2    // Type Inference
		const3, const4     = 3, 4 // Inline
		_                  = 5    // Unused
	)
	fmt.Println("Consts:", const1, const2, const3, const4)
}
