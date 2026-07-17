// Cmp
// The cmp package has functions for comparing two values of the same type.
// It also exports an type interface with all Ordered types (types that can be compared with <, >, <=, >=).

package library

import (
	"cmp"
	"fmt"
)

// CMP Functions
// The cmp package has functions for comparing two values of the same type.
func CmpFunctions() {

	// Compare
	// The compare function compares two values of the same type and returns -1, 0 or 1.
	x := cmp.Compare(1, 2)
	fmt.Println("x:", x) // Output: x: -1

	// Less
	// Less reports whether x is less than y.
	// For floating-point types, a NaN is considered less than any non-NaN, and -0.0 is not less
	// than (is equal to) 0.0.
	y := cmp.Less(1, 2)
	fmt.Println("y:", y) // Output: y: true

	// Or
	// Or returns the first of its arguments that is not equal to the zero value.
	// If no argument is non-zero, it returns the zero value.
	z := cmp.Or(0, 1, 2)
	fmt.Println("z:", z) // Output: z: 1
}

// Ordered
// The Ordered interface is the union of all types that can be compared with <, >, <=, >=.
func Compare[T cmp.Ordered](a, b T) int {
	return cmp.Compare(a, b)
}
