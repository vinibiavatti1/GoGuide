// Slices
// Slices are used to store a sequence of elements of the same type.
// Unlike arrays, slices can be resized dynamically.
// Syntax:
//   []<datatype>{<values>}

package containers

import "fmt"

// Declaring Slices
// Slices are declared without specifying a length.
// If a length is specified, an array will be created instead.
func DeclaringSlices() {
	// Declaring a Slice
	// When no elements are specified, the slice will be empty.
	// Output: x: []
	x := []int{}
	fmt.Println("x:", x)

	// Declaring a Slice with Predefined Values
	// We can initialize the slice with predefined values.
	// Output: [1 2 3]
	x = []int{1, 2, 3}
	fmt.Println("x:", x)

	// Declaring a Slice Specifying Values by Index (:)
	// We can specify values by index during the initialization by using the ":" operator.
	// Output: x: [1 0 3]
	x = []int{0: 1, 2: 3}
	fmt.Println("x:", x)

	// Declaring a Slice with Predefined Length
	// For this, we can use the "make()" function.
	// Output: [0 0 0]
	x = make([]int, 3)
	fmt.Println("x:", x)

	// Declaring a Slice with Predefined Length and Capacity
	// The capacity is the maximum number of elements that can be stored in the slice.
	// When the capacity is reached, the slice will be resized.
	// If the capacity is not specified, it will be equal to the length.
	// Output: [0 0 0] (length = 3, capacity = 5)
	x = make([]int, 3, 5)
	fmt.Println("x:", x)

	// Declaring a Slice of "any" Type
	// We can declare a slice of type "any" to store values of any type.
	// The "any" type is equivalent to the "interface{}" type in Go.
	// Output: y: [1 Hello true]
	y := []any{1, "Hello", true}
	fmt.Println("y:", y)
}

// Manipulating Slices
// We can add, access, reset and clear elements in a slice.
func ManipulatingSlices() {
	// Declaring a Slice
	// We will declare a slice with 3 elements to be used in the examples.
	x := []string{"A", "B", "C"}

	// Accessing Elements
	// We can access elements using indexes [i].
	// Output: y: B
	y := x[1]
	fmt.Println("y:", y)

	// Slicing Slices
	// We can slice slices using the [start:end] syntax.
	// Note: The end index is exclusive (not included in the slice).
	// Syntax: x[<start>:<end>]
	fmt.Println("x[ :2]", x[:2])  // Output: [A B]
	fmt.Println("x[1: ]", x[1:])  // Output: [B C]
	fmt.Println("x[1:2]", x[1:2]) // Output: [B]
	fmt.Println("x[ : ]", x[:])   // Output: [A B C] (Same reference)
	fmt.Println("x[0:0]", x[0:0]) // Output: [] (empty array)

	// Slicing with Capacity
	// We can specify the capacity of the slice using the capacity syntax.
	// Syntax: x[<start>:<end>:<capacity>]
	// Output: [A B] (length = 2, capacity = 5)
	// Output: [B] (length = 1, capacity = 5)
	fmt.Println("x[ :2:5]", x[:2:5])
	fmt.Println("x[1:2:5]", x[1:2:5])

	// Retrieving Length
	// We can get the length of a slice using the "len()" builtin function.
	// Output: Len: 3
	l := len(x)
	fmt.Println("Len:", l)

	// Mutating Elements
	// We can mutate elements using indexes [i].
	// Output: x: [A Z C]
	x[1] = "Z"
	fmt.Println("x:", x)

	// Appending Elements
	// We can use the "append()" builtin function to add elements to a slice.
	// Output: x: [A Z C D]
	x = append(x, "D")
	fmt.Println("x:", x)

	// Appending Multiple Elements
	// The "append()" function accepts a variadic number of elements.
	// Output: [A Z C D E F G]
	x = append(x, "E", "F", "G")
	fmt.Println("x:", x)

	// Resetting Elements
	// We can use the "clear" function to reset all elements of a slice to its default value.
	// Output: a: [0 0 0]
	a := []int{1, 2, 3}
	clear(a)
	fmt.Println("a:", a)

	// Clearing Slice
	// We can clear all elements from a slice by setting the slice using an empty range.
	// Output: x: []
	a = a[0:0]
	fmt.Println("a:", a)

	// Unpacking Slices
	// We can use the spread operator "..." to unpack slices into functions that accept a variable number of arguments.
	// Output: a: [1 2 3]
	a = append(a, []int{1, 2, 3}...)
	fmt.Println("a:", a)

	// Iterating Slices (Using standard for)
	// We can use the standard for loop to iterate over a slice
	// Output: 0 A | 1 Z | 2 C | 3 D | 4 E | 5 F | 6 G
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	// Iterating Slices (Using for-range)
	// We can use the for-range loop to iterate over a slice
	// Output: 0 A | 1 Z | 2 C | 3 D | 4 E | 5 F | 6 G
	for i, v := range x {
		fmt.Println(i, v)
	}
}
