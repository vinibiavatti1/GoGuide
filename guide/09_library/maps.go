// Maps
// Maps is a package that provides a set of functions to work with maps.
// It includes functions to create, update, delete, and retrieve values from maps.

package library

import (
	"fmt"
	"maps"
)

// Map Data Functions
// The following functions manipulate the data in maps.
func MapDataFunctions() {

	// Insert
	// Insert adds the key-value pairs from seq to m.
	// If a key in seq already exists in m, its value will be overwritten.
	x := map[string]int{"A": 1, "B": 2}
	y := map[string]int{"B": 8, "C": 9}
	maps.Insert(x, maps.All(y))
	fmt.Println(x) // Output: map[A:1 B:8 C:9]

	// Copy
	// Copy copies the key-value pairs from src to dst.
	// If a key in src already exists in dst, its value will be overwritten.
	// Same as Insert, but uses a map, not a sequence.
	x = map[string]int{"A": 1, "B": 2}
	y = map[string]int{"B": 8, "C": 9}
	maps.Copy(x, y)
	fmt.Println(x) // Output: map[A:1 B:8 C:9]

	// Clone
	// Clone returns a copy of m.
	// This is a shallow clone: the new keys and values are set using ordinary assignment.
	x = map[string]int{"A": 1, "B": 2, "C": 3}
	y = maps.Clone(x)
	fmt.Println(y) // Output: map[A:1 B:2 C:3]

	// DeleteFunc
	// DeleteFunc deletes any key/value pairs from m for which del returns true.
	x = map[string]int{"A": 1, "B": 2, "C": 3}
	maps.DeleteFunc(x, func(k string, v int) bool {
		return k == "B" || v == 3
	})
	fmt.Println(x) // Output: map[A:1]
}

// Map Compare Functions
// The following functions compare maps.
func MapCompareFunctions() {

	// Equal
	// Equal reports whether two maps contain the same key/value pairs. Values are compared using ==.
	x := map[string]int{"A": 1, "B": 2, "C": 3}
	eq := maps.Equal(x, map[string]int{"A": 1, "B": 2, "C": 3})
	fmt.Println(eq) // Output: true

	// EqualFunc
	// EqualFunc reports whether two maps contain the same key/value pairs.
	// Values are compared using the provided function.
	x = map[string]int{"A": 1, "B": 2, "C": 3}
	eq = maps.EqualFunc(x, map[string]int{"A": 2, "B": 3, "C": 4}, func(a, b int) bool {
		return a == b+1
	})
	fmt.Println(eq) // Output: true
}

// Map Seq Functions
// The following functions convert between maps and iterators.
func MapSeqFunctions() {

	// All
	// All returns an iterator over key-value pairs from m.
	// The iteration order is not specified and is not guaranteed to be the same from one call to the next.
	x := map[string]int{"A": 1, "B": 2, "C": 3}
	for k, y := range maps.All(x) {
		fmt.Println(k, y) // Output: A 1, B 2, C 3
	}

	// Keys
	// Keys returns an iterator over keys in m.
	// The iteration order is not specified and is not guaranteed to be the same from one call to the next.
	x = map[string]int{"A": 1, "B": 2, "C": 3}
	for k := range maps.Keys(x) {
		fmt.Println(k) // Output: A, B, C
	}

	// Values
	// Values returns an iterator over values in m.
	// The iteration order is not specified and is not guaranteed to be the same from one call to the next.
	x = map[string]int{"A": 1, "B": 2, "C": 3}
	for v := range maps.Values(x) {
		fmt.Println(v) // Output: 1, 2, 3
	}

	// Collect
	// Collect collects key-value pairs from seq into a new map and returns it.
	x = map[string]int{"A": 1, "B": 2, "C": 3}
	y := maps.Collect(maps.All(x))
	fmt.Println(y) // Output: map[A:1 B:2 C:3]
}
