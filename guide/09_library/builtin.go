// Builtin
// The builtin package provides a set of built-in functions for common operations.
// The functions are available in all Go programs without the need for an import statement.

package library

import "fmt"

// Builtin Functions
// The function below shows examples of all builtin functions.
func BuiltinFunctions() {

	// Append
	// The append built-in function appends elements to the end of a slice.
	// Signature: append(slice []Type, elems ...Type) []Type
	slc := []string{"A", "B"}
	slc = append(slc, "C", "D")
	fmt.Println(slc) // Output: [A B C D]

	// Cap
	// The cap built-in function returns the capacity of v, according to its type:
	// - Array: the number of elements in v (same as len(v)).
	// - Pointer to array: the number of elements in *v (same as len(v)).
	// - Slice: the maximum length the slice can reach when resliced; if v is nil, cap(v) is zero.
	// - Channel: the channel buffer capacity, in units of elements; if v is nil, cap(v) is zero.
	// Signature: cap(v Type) int
	slc = []string{"A", "B", "C"}
	capacity := cap(slc)
	fmt.Println(capacity) // Output: 3

	// Clear
	// The clear built-in function clears maps and slices.
	// - For maps, clear deletes all entries, resulting in an empty map.
	// - For slices, resets all elements to their zero value.
	// Signature: clear[T ~[]Type | ~map[Type]Type1](t T)
	slc2 := []int{1, 2, 3}
	clear(slc2)
	fmt.Println(slc2) // Output: [0 0 0]

	// Close
	// The close built-in function closes a channel, which must be either bidirectional or send-only.
	// It should be executed only by the sender, never the receiver.
	// Signature: close(c chan<- Type)
	ch := make(chan int)
	close(ch)

	// Complex
	// The complex built-in function constructs a complex value from two floating-point values.
	// Signature: complex(r, i FloatType) ComplexType
	c := complex(1, 2)
	fmt.Println(c) // Output: (1+2i)

	// Copy
	// The copy built-in function copies elements from a source slice into a destination slice.
	// Signature: copy(dst, src []Type) int
	src := []int{1, 2, 3}
	dst := make([]int, 3)
	count := copy(dst, src)
	fmt.Println(count, dst) // Output: 3 [1 2 3]

	// Delete
	// The delete built-in function deletes the element with the specified key from the map
	// Signature: delete(m map[Type]Type1, key Type)
	mp := map[string]int{
		"A": 1,
		"B": 2,
	}
	delete(mp, "A")
	fmt.Println(mp) // Output: map[B:2]

	// Imag
	// The imag built-in function returns the imaginary part of the complex number.
	// Signature: imag(c ComplexType) FloatType
	c = 1 + 2i
	imagPart := imag(c)
	fmt.Println(imagPart) // Output: 2

	// Len
	// The len built-in function returns the length of v, according to its type:
	// - Array: the number of elements in v.
	// - Pointer to array: the number of elements in *v (even if v is nil).
	// - Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
	// - String: the number of bytes in v.
	// - Channel: the number of elements queued (unread) in the channel buffer; if v is nil, len(v) is zero.
	// Signature: len(v Type) int
	slc = []string{"A", "B", "C"}
	length := len(slc)
	fmt.Println(length) // Output: 3

	// Make
	// The make built-in function allocates and initializes an object of type slice, map, or chan (only).
	// Unlike new, make's return type is the same as the type of its argument, not a pointer to it.
	// Signature: make(t Type, size ...IntegerType) Type
	slc2 = make([]int, 3)
	fmt.Println(slc2) // Output: [0 0 0]

	// Max
	// The max built-in function returns the largest value of a fixed number of arguments of cmp.Ordered types.
	// Signature: max[T cmp.Ordered](x T, y ...T) T
	maxValue := max(1, 2, 3, 4)
	fmt.Println(maxValue) // Output: 4

	// Min
	// The min built-in function returns the smallest value of a fixed number of arguments of cmp.Ordered types
	// Signature: min[T cmp.Ordered](x T, y ...T) T
	minValue := min(1, 2, 3, 4)
	fmt.Println(minValue) // Output: 1

	// New
	// The new built-in function allocates memory.
	// The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.
	// Signature: new(Type) *Type
	ptr := new(int)
	fmt.Println(*ptr) // Output: 0

	// Panic
	// The panic built-in function stops normal execution of the current goroutine.
	// When a function F calls panic, normal execution of F stops immediately.
	// Signature: panic(v any)
	_ = func() {
		panic("Something went wrong")
	}

	// Print
	// The print built-in function formats its arguments in an implementation-specific way and writes the result to standard error.
	// Note: Print is useful for bootstrapping and debugging. It is preferred to use fmt package for print operations.
	// Signature: print(args ...Type)
	print("Hello, World!")

	// PrintLn
	// Similar to print, but appends a newline character at the end.
	// Note: PrintLn is useful for bootstrapping and debugging. It is preferred to use fmt package for print operations.
	// Signature: println(args ...Type)
	println("Hello, World!")

	// Real
	// The real built-in function returns the real part of the complex number.
	// Signature: real(c ComplexType) FloatType
	c = 1 + 2i
	realPart := real(c)
	fmt.Println(realPart) // Output: 1

	// Recover
	// The recover built-in function allows a program to manage behavior of a panicking goroutine.
	// Executing a call to recover inside a deferred function (but not any function called by it) stops the
	// panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic.
	// Signature: recover() any
	_ = func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		panic("Something went wrong")
	}
}
