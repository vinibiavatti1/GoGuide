// Slices
// The slices package provides a set of functions for manipulating slices.
// It includes functions for filtering, mapping, and reducing slices.
// It also includes functions for finding the minimum and maximum values in a slice.

package library

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

// Slice Data Functions
// The following functions manipulate the data in slices.
func SliceDataFunctions() {

	// Insert
	// Insert inserts the values v... into s at index i, returning the modified slice.
	x := []string{"A", "B", "C"}
	x = slices.Insert(x, 1, "X", "Y")
	fmt.Println(x) // Output: [A X Y B C] (X and Y are inserted at index 1)

	// Concat
	// Concat returns a new slice concatenating the passed in slices.
	x = []string{"A", "B", "C"}
	x = slices.Concat(x, []string{"D", "E", "F"})
	fmt.Println(x) // Output: [A B C D E F]

	// Delete
	// Delete removes the elements s[i:j] from s, returning the modified slice
	x = []string{"A", "B", "C", "D", "E"}
	x = slices.Delete(x, 1, 3)
	fmt.Println(x) // Output: [A D E] (B and C are removed)

	// DeleteFunc
	// DeleteFunc works like Delete, but uses a custom comparison function.
	x = []string{"A", "B", "C"}
	x = slices.DeleteFunc(x, func(a string) bool {
		return a == "B"
	})
	fmt.Println(x) // Output: [A C] (B is removed)

	// Repeat
	// Repeat returns a new slice that repeats the provided slice the given number of times.
	x = []string{"A", "B", "C"}
	x = slices.Repeat(x, 2)
	fmt.Println(x) // Output: [A B C A B C] (repeated twice)

	// Replace
	// Replace replaces the elements s[i:j] by the given v, and returns the modified slice.
	x = []string{"A", "B", "C"}
	x = slices.Replace(x, 1, 2, "X")
	fmt.Println(x) // Output: [A X C] (B is replaced by X)

	// Compact
	// Compact replaces consecutive runs of equal elements with a single copy.
	// This is like the uniq command found on Unix.
	// Compact modifies the contents of the slice s and returns the modified slice, which may have a smaller length.
	// Compact zeroes the elements between the new length and the original length.
	x = []string{"A", "B", "B", "C", "C", "C"}
	x = slices.Compact(x)
	fmt.Println(x) // Output: [A B C C C] (length is 3, capacity is 6)

	// CompactFunc
	// CompactFunc works like Compact, but uses a custom comparison function.
	x = []string{"A", "b", "B", "C", "c", "C"}
	x = slices.CompactFunc(x, strings.EqualFold)
	fmt.Println(x) // Output: [A b C C C] (length is 3, capacity is 6)

	// Reverse
	// Reverse reverses the elements of the slice in place.
	// Note: It modifies the original slice.
	x = []string{"A", "B", "C"}
	slices.Reverse(x)
	fmt.Println(x) // Output: [C B A] (slices are reversed)

	// AppendSeq
	// AppendSeq appends the values from seq to the slice and returns the extended slice.
	x = []string{"A", "B", "C"}
	x = slices.AppendSeq(x, slices.Values(x))
	fmt.Println(x) // Output: [A B C A B C]

	// Clone
	// Clone returns a copy of the slice. The elements are copied using assignment, so this is a shallow clone.
	// The result may have additional unused capacity.
	x = []string{"A", "B", "C"}
	clone := slices.Clone(x)
	fmt.Println(clone) // Output: [A B C]
}

// Slice Index Functions
// The following functions search for indexes in slices.
func SliceIndexFunctions() {

	// Index
	// Index returns the index of the first occurrence of v in s, or -1 if not present.
	x := []string{"A", "B", "C"}
	i := slices.Index(x, "B")
	fmt.Println(i) // Output: 1 (B is at index 1)

	// IndexFunc
	// IndexFunc returns the first index i satisfying f(s[i]), or -1 if none do.
	x = []string{"A", "b", "C"}
	i = slices.IndexFunc(x, func(a string) bool {
		return strings.ToUpper(a) == "B"
	})
	fmt.Println(i) // Output: 1 (B is at index 1)
}

// Slice Search Functions
// The following functions search for elements in slices.
func SliceSearchFunctions() {

	// Binary Seach
	// BinarySearch searches for target in a sorted slice and returns the earliest position where target is found,
	// or the position where target would appear in the sort order.
	// It also returns a bool saying whether the target is really found in the slice.
	// The slice must be sorted in increasing order.
	x := []string{"A", "B", "C"}
	i, found := slices.BinarySearch(x, "B")
	fmt.Println(i, found) // Output: 1 true

	// Binary Search Func
	// BinarySearchFunc works like BinarySearch, but uses a custom comparison function.
	// The comparison function must return a comparison result: (-1, 0 or 1).
	type Person struct {
		ID   int
		Name string
	}
	y := []*Person{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Maria"},
	}
	i, found = slices.BinarySearchFunc(y, &Person{Name: "Maria"}, func(a *Person, b *Person) int {
		return cmp.Compare(a.Name, b.Name)
	})
	fmt.Println(i, found) // Output: 1 true

	// Contains
	// Contains reports whether v is present in s.
	x = []string{"A", "B", "C"}
	has := slices.Contains(x, "B")
	fmt.Println(has) // Output: true (B is present)

	// ContainsFunc
	// ContainsFunc works like Contains, but uses a custom comparison function.
	x = []string{"A", "B", "C"}
	has = slices.ContainsFunc(x, func(a string) bool {
		return a == "B"
	})
	fmt.Println(has) // Output: true (B is present)

	// Max
	// Max returns the maximal value in x.
	z := []int{29, 53, 42}
	max := slices.Max(z)
	fmt.Println(max) // Output: 53 (maximal value is 53)

	// MaxFunc
	// MaxFunc works like Max, but uses a custom comparison function.
	z = []int{29, 53, 42}
	max = slices.MaxFunc(z, func(a int, b int) int {
		return a - b
	})
	fmt.Println(max) // Output: 53 (maximal value is 53)

	// Min
	// Min returns the minimal value in x.
	min := slices.Min(z)
	fmt.Println(min) // Output: 29 (minimal value is 29)

	// MinFunc
	// MinFunc works like Min, but uses a custom comparison function.
	min = slices.MinFunc(z, func(a int, b int) int {
		return a - b
	})
	fmt.Println(min) // Output: 29 (minimal value is 29)
}

// Slice Compare Functions
// The following functions compare slices.
func SliceCompareFunctions() {

	// Compare
	// Compare compares the elements of s1 and s2, using cmp.Compare on each pair of elements.
	x := []string{"A", "B", "C"}
	cmp := slices.Compare(x, []string{"A", "B", "C"})
	fmt.Println(cmp) // Output: 0 (equal)

	// CompareFunc
	// CompareFunc works like Compare, but uses a custom comparison function.
	x = []string{"A", "B", "C"}
	cmp = slices.CompareFunc(x, []string{"a", "b", "c"}, func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(cmp) // Output: 0 (equal)

	// Equal
	// Equal reports whether two slices are equal.
	x = []string{"A", "B", "C"}
	eq := slices.Equal(x, []string{"A", "B", "C"})
	fmt.Println(eq) // Output: true (slices are equal)

	// EqualFunc
	// EqualFunc works like Equal, but uses a custom comparison function.
	x = []string{"A", "B", "C"}
	eq = slices.EqualFunc(x, []string{"a", "b", "c"}, strings.EqualFold)
	fmt.Println(eq) // Output: true (slices are equal)
}

// Slice Sort Functions
// The following functions sort slices in various ways.
func SliceSortFunctions() {

	// Sort
	// Sort sorts a slice of any ordered type in ascending order.
	// Note: It modifies the original slice.
	x := []string{"C", "A", "B"}
	slices.Sort(x)
	fmt.Println(x) // Output: [A B C] (slices are sorted)

	// SortFunc
	// SortFunc sorts a slice of any type using the provided comparison function.
	// Note: It modifies the original slice.
	x = []string{"C", "A", "b"}
	slices.SortFunc(x, func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(x) // Output: [A b C] (slices are sorted)

	// SortStableFunc
	// SortStableFunc sorts the slice x while keeping the original order of equal elements,
	// using cmp to compare elements in the same way as SortFunc.
	// Note: It modifies the original slice.
	x = []string{"C", "A", "b"}
	slices.SortStableFunc(x, func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(x) // Output: [A b C] (slices are sorted)

	// Sorted
	// Sorted collects values from seq into a new slice, sorts the slice, and returns it.
	x = []string{"C", "A", "b"}
	sorted := slices.Sorted(slices.Values(x))
	fmt.Println(sorted) // Output: [A b C] (slices are sorted)

	// SortedFunc
	// SortedFunc collects values from seq into a new slice, sorts the slice using the comparison function, and returns it.
	x = []string{"C", "A", "b"}
	sorted = slices.SortedFunc(slices.Values(x), func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(sorted) // Output: [A b C] (slices are sorted)

	// SortedStableFunc
	// SortedStableFunc collects values from seq into a new slice.
	// It then sorts the slice while keeping the original order of equal elements,
	// using the comparison function to compare elements.
	x = []string{"C", "A", "b"}
	sorted = slices.SortedStableFunc(slices.Values(x), func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(sorted) // Output: [A b C] (slices are sorted)

	// IsSorted
	// IsSorted reports whether x is sorted in ascending order.
	x = []string{"A", "B", "C"}
	is := slices.IsSorted(x)
	fmt.Println(is) // Output: true (slices are sorted)

	// IsSortedFunc
	// IsSortedFunc works like IsSorted, but uses a custom comparison function.
	x = []string{"A", "b", "C", "d"}
	is = slices.IsSortedFunc(x, func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(is) // Output: true (slices are sorted)
}

// Slice Seq Functions
// The following functions convert between slices and iterators.
func SliceSeqFunctions() {

	// All (slice -> Seq2)
	// All returns an iterator over index-value pairs in the slice in the usual order.
	x := []string{"A", "B", "C"}
	for i, v := range slices.All(x) {
		fmt.Println(i, v) // Output: 0 A, 1 B, 2 C
	}

	// Values (slice -> Seq)
	// Values returns an iterator that yields the slice elements in order.
	x = []string{"A", "B", "C"}
	for v := range slices.Values(x) {
		fmt.Println(v) // Output: A B C
	}

	// Collect (Seq -> slice)
	// Collect collects values from seq into a new slice and returns it.
	x = []string{"A", "B", "C"}
	y := slices.Collect(slices.Values(x))
	fmt.Println(y) // Output: [A B C]
}

// Slice Other Functions
// The following functions are other available functions in the slices package.
func SliceOtherFunctions() {

	// Backward
	// Backward returns an iterator over index-value pairs in the slice, traversing it backward with descending indices.
	x := []string{"A", "B", "C"}
	for i, v := range slices.Backward(x) {
		fmt.Println(i, v) // Output: 2 C, 1 B, 0 A
	}

	// Chunk
	// Chunk returns an iterator over consecutive sub-slices of up to n elements of s.
	// All but the last sub-slice will have size n. All sub-slices are clipped to have no capacity beyond the length.
	// If s is empty, the sequence is empty: there is no empty slice in the sequence. Chunk panics if n is less than 1.
	x = []string{"A", "B", "C", "D", "E"}
	chunk := slices.Chunk(x, 2)
	for v := range chunk {
		fmt.Println(v) // Output: [A B], [C D], [E]
	}

	// Clip
	// Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
	x = make([]string, 0, 10) // (capacity is 10)
	x = append(x, "A", "B", "C")
	x = slices.Clip(x)
	fmt.Println(cap(x)) // Output: 3 (capacity is 3)

	// Grow
	// Grow increases the slice's capacity, if necessary, to guarantee space for another n elements.
	x = []string{"A", "B", "C"}
	x = slices.Grow(x, 2)
	fmt.Println(cap(x)) // Output: 5 (capacity is 5)
}
