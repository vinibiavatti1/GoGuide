// Generics
// - Generics allow you to write functions and data structures that can operate on different types
//   without sacrificing type safety. This is done using type parameters.
// - Generics are a powerful feature introduced in Go 1.18, allowing developers to write more
//   flexible and reusable code.
// - Syntax (Definition):
//   - func <name>[T <types>](<params> T) T {...}
//   - struct <name>[T <types>] {...}
// - Syntax (Usage):
//   - <name>[<type>](<params> T) T {...}
//   - <name>[<type>] {...}

package typing

import (
	"cmp"
	"fmt"
)

// Generic Function
// - The function below demonstrates a generic function that takes two parameters of the same type
//   and returns their sum.
// - The type parameter T can be int or float64, allowing the function to work with both integer and
//   floating-point numbers.

func sumNumbers[T int | float64](x, y T) T {
	return x + y
}

// Perform Generic Function
// - Now, we will use the function to add two integers and two floats.

func PerformSumNumbers() {

	// Sum two integers
	// - The generic function works to integer types.
	// - Note that the type parameter T is inferred from the arguments passed to the function.
	// - Output: Integer Sum: 30

	intSum := sumNumbers(10, 20)
	println("Integer Sum:", intSum)

	// Sum two floats
	// - The generic function works to float types.
	// - Output: Float Sum: 30.8

	floatSum := sumNumbers(10.5, 20.3)
	println("Float Sum:", floatSum)
}

//==================================================================================================
// Type Approximation "~"
//==================================================================================================

// Type Approximation "~"
// - Type approximation is a token that can be used in types to indicate that types based on the
//   original type are allowed as well, not only the original one.
// - For example, we will define a nwe type based on int, but it will not be available for the
//   generic function, since the function only accepts the original type itself, not approximations
//   of it.

type Integer int

// Defining Generic Function with Type Approximation
// - Now, to allow the function to accept related types, we will define the function with the type
//   approximation token "~".

func sumAnyNumbers[T ~int | ~float64](x, y T) T {
	return x + y
}

// Perform Generic Function with Type Approximation
// - Now, we will use the function to add two integers and two floats using the type approximation.

func PerformSumAnyNumbers() {

	// Sum two Integers
	// - The generic function works to Integer types (type based on int).
	// - Note that the type parameter T is inferred from the arguments passed to the function.
	// - Integer Sum: Output: 30

	var x Integer = 10
	var y Integer = 20
	intSum := sumAnyNumbers(x, y)
	println("Integer Sum:", intSum)
}

//==================================================================================================
// Type Interface
//==================================================================================================

// Type Interface
// - We can define a type interface that represents a set of types that can be used with the generic
//   function.
// - This allows us to create a group of types.
// - We can use the "|" operator to separate the types.

type Numeric interface {
	~int | ~float64
}

// Generic Function with Type Interface
// - Now, instead of setting all the possible types for the function parameter, we can use the
//   Numeric interface instead.
// - In this case, "T" can be of ~int or ~float64 type, and the function will work with both types.

func sumNumerics[T Numeric](x, y T) T {
	return x + y
}

// Perform Generic Function with Type Interface
// - Now, we will use the function to add two integers and two floats using the Numeric interface.
// - Perform Generic Function
// - Now, we will use the function to add two integers and two floats.

func PerformSumNumerics() {

	// Sum two integers
	// - The generic function works to integer types.
	// - Note that the type parameter T is inferred from the arguments passed to the function.
	// - Output: 30

	intSum := sumNumerics(10, 20)
	println("Integer Sum:", intSum)

	// Sum two floats
	// - The generic function works to float types.
	// - Output: 30.8

	floatSum := sumNumerics(10.5, 20.3)
	println("Float Sum:", floatSum)
}

//==================================================================================================
// Generic Structs & Interfaces
//==================================================================================================

// Generic Struct
// - We can also define a generic struct that can hold values of different types.
// - This allows us to create data structures that can work with any type.

type Pair[K any, T any] struct {
	Key   K
	Value T
}

// Generic Interface
// - We can define a generic interface that can be implemented by different types.
// - This allows us to create a common interface for different types and use it in the generic
//   function.

type Getter[T any] interface {
	GetValue() T
}

// Implementing the Getter interface for the Pair struct
// - The Pair struct implements the Getter interface, allowing us to use the GetValue method.

func (p *Pair[K, V]) GetValue() V {
	return p.Value
}

// Using Generic Struct
// - Now, we will create a generic struct instance and use it to hold different types of values.
// - The type parameter T can be any type, allowing the struct to hold values of different types.

func UsingGenericStruct() {

	// Creating a Generic Struct of Int
	// - We will create a generic struct instance that holds an integer value.
	// - Output: x: {key 10}

	x := Pair[string, int]{
		Key:   "key",
		Value: 10,
	}
	fmt.Println("x:", x)

	// Creating a Generic Struct of String
	// - We will create a generic struct instance that holds a string value.
	// - Output: y: {10 value}

	y := Pair[int, string]{
		Key:   10,
		Value: "value",
	}
	fmt.Println("y:", y)

	// Using the Getter interface
	// - We can use the Getter interface to get the value from the generic struct.
	// - Output: 10 value

	fmt.Println(x.GetValue(), y.GetValue())
}

//==================================================================================================
// Generic Containers
//==================================================================================================

// Generic Containers
// - In this example, we will create a map of numbers and use the generic function to sum the values
//   in the map.
// - The "cmp.Ordered" constraint is used to ensure that the values in the map are ordered types
//   (int, float, etc.).

func sumMap[K comparable, V cmp.Ordered](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Perform Generic Containers
// - Now, we will use the function to sum the values in the map.

func PerformSumMap() {

	// Initialize a map for the integer values
	// - The map keys are strings and the values are integers.
	// - Output: Integer map: 46

	x := map[string]int{
		"first":  34,
		"second": 12,
	}
	fmt.Println("Integer map:", sumMap(x))

	// Initialize a map for the float values
	// - The map keys are strings and the values are floats.
	// - Output: Float map: 62.97

	y := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Println("Float map:", sumMap(y))
}
