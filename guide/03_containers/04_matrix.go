// Matrix
// A matrix is a two-dimensional array. We can use a slice of slices to represent a matrix in Go.
// We can have arrays and slices with n dimensions
// The examples below will only show 2D arrays/slices to demonstrate the concept of multidimensional arrays/slices.
// Syntax:
// - 2D: [][]<datatype>{}
// - 3D: [][][]<datatype>{}

package containers

import "fmt"

// Matrices
// A matrix is a two-dimensional array. We can use a slice of slices or an array of arrays to represent a
// matrix in Go.
func DeclaringMatrices() {
	// Declaring a Matrix with Arrays
	// The example below shows how to declare an array of arrays.
	// Output: x: [[A B C] [D E F] [G H I]]
	x := [3][3]string{
		{"A", "B", "C"},
		{"D", "E", "F"},
		{"G", "H", "I"},
	}
	fmt.Println("x:", x)

	// Declaring a Matrix with Slices
	// We can use slices of slices to represent a matrix in Go as well.
	// Output: y: [[A B C] [D E F] [G H I]]
	y := [][]string{
		{"A", "B", "C"},
		{"D", "E", "F"},
		{"G", "H", "I"},
	}
	fmt.Println("y:", y)
}

// Manipulating Matrices
// We can access, mutate and iterate a matrix.
func ManipulatingMatrices() {
	// Declaring a Matrix
	// We will declare a matrix to be used in the examples.
	x := [3][3]string{
		{"A", "B", "C"},
		{"D", "E", "F"},
		{"G", "H", "I"},
	}

	// Accessing Elements
	// We can access elements using indexes [i][j].
	// Output: y: E
	y := x[1][1]
	fmt.Println("y:", y)

	// Retrieving Length
	// We can get the length of the matrix using the "len()" builtin function.
	// Output: Rows: 3 Cols: 3
	lrow := len(x)
	lcol := len(x[0])
	fmt.Println("Rows:", lrow, "Cols:", lcol)

	// Mutating Elements
	// We can mutate elements using indexes [i].
	// Note: Arrays are fixed in size, so we cannot append or remove elements.
	// Output: x: [[A B C] [D Z F] [G H I]]
	x[1][1] = "Z"
	fmt.Println("x:", x)

	// Iterating Matrices (Using standard for)
	// We can use the standard for loop to iterate over an array.
	// Output: 0 0 A | 0 1 B | 0 2 C | 1 0 D | 1 1 Z | 1 2 F | 2 0 G | 2 1 H | 2 2 I
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			fmt.Println(i, j, x[i][j])
		}
	}

	// Iterating Matrices (Using for-range)
	// We can use the for-range loop to iterate over an array.
	// Output: 0 0 A | 0 1 B | 0 2 C | 1 0 D | 1 1 Z | 1 2 F | 2 0 G | 2 1 H | 2 2 I
	for i, row := range x {
		for j, v := range row {
			fmt.Println(i, j, v)
		}
	}
}
