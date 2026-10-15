// Slices
// - The slices package provides a set of functions for manipulating slices.
// - It includes functions for filtering, mapping, and reducing slices.
// - It also includes functions for finding the minimum and maximum values in a slice.

package library

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

//==================================================================================================
// Element Functions
//==================================================================================================

func SliceElementFunctions() {

	// Insert
	// - Insert inserts the values v... into s at index i, returning the modified slice.
	// - Output: [A X Y B C] (X and Y are inserted at index 1)

	x := []string{"A", "B", "C"}
	x = slices.Insert(x, 1, "X", "Y")
	fmt.Println(x)

	// Concat
	// - Concat returns a new slice concatenating the passed in slices.
	// - Output: [A B C D E F]

	x = []string{"A", "B", "C"}
	x = slices.Concat(x, []string{"D", "E", "F"})
	fmt.Println(x)

	// Delete
	// - Delete removes the elements s[i:j] from s, returning the modified slice
	// - Output: [A D E] (B and C are removed)

	x = []string{"A", "B", "C", "D", "E"}
	x = slices.Delete(x, 1, 3)
	fmt.Println(x)

	// DeleteFunc
	// - DeleteFunc works like Delete, but uses a custom comparison function.
	// - Output: [A C] (B is removed)

	x = []string{"A", "B", "C"}
	x = slices.DeleteFunc(x, func(a string) bool {
		return a == "B"
	})
	fmt.Println(x)

	// Repeat
	// - Repeat returns a new slice that repeats the provided slice the given number of times.
	// - Output: [A B C A B C]

	x = []string{"A", "B", "C"}
	x = slices.Repeat(x, 2)
	fmt.Println(x)

	// Replace
	// - Replace replaces the elements s[i:j] by the given v, and returns the modified slice.
	// - Output: [A X X C]

	x = []string{"A", "B", "C", "D"}
	x = slices.Replace(x, 1, 3, "X")
	fmt.Println(x)

	// Compact
	// - Compact replaces consecutive runs of equal elements with a single copy.
	// - This is like the uniq command found on Unix.
	// - Compact modifies the contents of the slice s and returns the modified slice, which may have
	//   a smaller length.
	// - Compact zeroes the elements between the new length and the original length.
	// - Output: [A B C] (length is 3, capacity is 6)

	x = []string{"A", "A", "B", "B", "C", "C"}
	x = slices.Compact(x)
	fmt.Println(x)

	// CompactFunc
	// - CompactFunc works like Compact, but uses a custom comparison function.
	// - Output: [A b C] (length is 3, capacity is 6)

	x = []string{"A", "a", "b", "B", "C", "c"}
	x = slices.CompactFunc(x, strings.EqualFold)
	fmt.Println(x)

	// Reverse
	// - Reverse reverses the elements of the slice in place.
	// - Note: It modifies the original slice.
	// - Output: [C B A] (slices are reversed)

	x = []string{"A", "B", "C"}
	slices.Reverse(x)
	fmt.Println(x)

	// AppendSeq
	// - AppendSeq appends the values from seq to the slice and returns the extended slice.
	// - Output: [A B C A B C]

	x = []string{"A", "B", "C"}
	x = slices.AppendSeq(x, slices.Values(x))
	fmt.Println(x)

	// Clone
	// - Clone returns a copy of the slice. The elements are copied using assignment, so this is a
	//   shallow clone.
	// - The result may have additional unused capacity.
	// - Output: [A B C]

	x = []string{"A", "B", "C"}
	clone := slices.Clone(x)
	fmt.Println(clone)
}

//==================================================================================================
// Index Functions
//==================================================================================================

func SliceIndexFunctions() {

	// Index
	// - Index returns the index of the first occurrence of v in s, or -1 if not present.
	// - Output: 1 (B is at index 1)

	x := []string{"A", "B", "C"}
	i := slices.Index(x, "B")
	fmt.Println(i)

	// IndexFunc
	// - IndexFunc returns the first index i satisfying f(s[i]), or -1 if none do.
	// - Output: 1 (B is at index 1)

	x = []string{"A", "b", "C"}
	i = slices.IndexFunc(x, func(a string) bool {
		return strings.ToUpper(a) == "B"
	})
	fmt.Println(i)
}

//==================================================================================================
// Search Functions
//==================================================================================================

func SliceSearchFunctions() {

	// Binary Seach
	// - BinarySearch searches for target in a sorted slice and returns the earliest position where
	//   target is found, or the position where target would appear in the sort order.
	// - It also returns a bool saying whether the target is really found in the slice.
	// - The slice must be sorted in increasing order.
	// - Output: 1 true

	x := []string{"A", "B", "C"}
	i, found := slices.BinarySearch(x, "B")
	fmt.Println(i, found)

	// Binary Search Func
	// - BinarySearchFunc works like BinarySearch, but uses a custom comparison function.
	// - The comparison function must return a comparison result: (-1, 0 or 1).
	// - Output: 1 true

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
	fmt.Println(i, found)

	// Contains
	// - Contains reports whether v is present in s.
	// - Output: true (B is present)

	x = []string{"A", "B", "C"}
	has := slices.Contains(x, "B")
	fmt.Println(has)

	// ContainsFunc
	// - ContainsFunc works like Contains, but uses a custom comparison function.
	// - Output: true (B is present)

	x = []string{"A", "B", "C"}
	has = slices.ContainsFunc(x, func(a string) bool {
		return a == "B"
	})
	fmt.Println(has)

	// Max
	// - Max returns the maximal value in x.
	// - Output: 53 (maximal value is 53)

	z := []int{29, 53, 42}
	max := slices.Max(z)
	fmt.Println(max)

	// MaxFunc
	// - MaxFunc works like Max, but uses a custom comparison function.
	// - Output: 53 (maximal value is 53)

	z = []int{29, 53, 42}
	max = slices.MaxFunc(z, func(a int, b int) int {
		return a - b
	})
	fmt.Println(max)

	// Min
	// - Min returns the minimal value in x.
	// - Output: 29 (minimal value is 29)

	min := slices.Min(z)
	fmt.Println(min)

	// MinFunc
	// - MinFunc works like Min, but uses a custom comparison function.
	// - Output: 29 (minimal value is 29)

	min = slices.MinFunc(z, func(a int, b int) int {
		return a - b
	})
	fmt.Println(min)
}

//==================================================================================================
// Compare Functions
//==================================================================================================

func SliceCompareFunctions() {

	// Compare
	// - Compare compares the elements of s1 and s2, using cmp.Compare on each pair of elements.
	// - Output: 0 (equal)

	x := []string{"A", "B", "C"}
	cmp := slices.Compare(x, []string{"A", "B", "C"})
	fmt.Println(cmp)

	// CompareFunc
	// - CompareFunc works like Compare, but uses a custom comparison function.
	// - Output: 0 (equal)

	x = []string{"A", "B", "C"}
	cmp = slices.CompareFunc(x, []string{"a", "b", "c"}, func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(cmp)

	// Equal
	// - Equal reports whether two slices are equal.
	// - Output: true (slices are equal)

	x = []string{"A", "B", "C"}
	eq := slices.Equal(x, []string{"A", "B", "C"})
	fmt.Println(eq)

	// EqualFunc
	// - EqualFunc works like Equal, but uses a custom comparison function.
	// - Output: true (slices are equal)

	x = []string{"A", "B", "C"}
	eq = slices.EqualFunc(x, []string{"a", "b", "c"}, strings.EqualFold)
	fmt.Println(eq)
}

//==================================================================================================
// Sort Functions
//==================================================================================================

func SliceSortFunctions() {

	// Sort
	// - Sort sorts a slice of any ordered type in ascending order.
	// - Note: It modifies the original slice.
	// - Output: [A B C] (slices are sorted)

	x := []string{"C", "A", "B"}
	slices.Sort(x)
	fmt.Println(x)

	// SortFunc
	// - SortFunc sorts a slice of any type using the provided comparison function.
	// - Note: It modifies the original slice.
	// - Output: [A b C] (slices are sorted)

	x = []string{"C", "A", "b"}
	slices.SortFunc(x, func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(x)

	// SortStableFunc
	// - SortStableFunc sorts the slice x while keeping the original order of equal elements, using
	//   cmp to compare elements in the same way as SortFunc.
	// - Note: It modifies the original slice.
	// - Output: [A b C] (slices are sorted)

	x = []string{"C", "A", "b"}
	slices.SortStableFunc(x, func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(x)

	// Sorted
	// - Sorted collects values from seq into a new slice, sorts the slice, and returns it.
	// - Output: [A b C] (slices are sorted)

	x = []string{"C", "A", "b"}
	sorted := slices.Sorted(slices.Values(x))
	fmt.Println(sorted)

	// SortedFunc
	// - SortedFunc collects values from seq into a new slice, sorts the slice using the comparison
	//   function, and returns it.
	// - Output: [A b C] (slices are sorted)

	x = []string{"C", "A", "b"}
	sorted = slices.SortedFunc(slices.Values(x), func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(sorted)

	// SortedStableFunc
	// - SortedStableFunc collects values from seq into a new slice.
	// - It then sorts the slice while keeping the original order of equal elements, using the
	//   comparison function to compare elements.
	// - Output: [A b C] (slices are sorted)

	x = []string{"C", "A", "b"}
	sorted = slices.SortedStableFunc(slices.Values(x), func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(sorted)

	// IsSorted
	// - IsSorted reports whether x is sorted in ascending order.
	// - Output: true (slices are sorted)

	x = []string{"A", "B", "C"}
	is := slices.IsSorted(x)
	fmt.Println(is)

	// IsSortedFunc
	// - IsSortedFunc works like IsSorted, but uses a custom comparison function.
	// - Output: true (slices are sorted)

	x = []string{"A", "b", "C", "d"}
	is = slices.IsSortedFunc(x, func(a string, b string) int {
		return strings.Compare(strings.ToUpper(a), strings.ToUpper(b))
	})
	fmt.Println(is)
}

//==================================================================================================
// Seq Functions
//==================================================================================================

func SliceSeqFunctions() {

	// All (slice -> Seq2)
	// - All returns an iterator over index-value pairs in the slice in the usual order.
	// - Output: 0 A, 1 B, 2 C

	x := []string{"A", "B", "C"}
	for i, v := range slices.All(x) {
		fmt.Println(i, v)
	}

	// Values (slice -> Seq)
	// - Values returns an iterator that yields the slice elements in order.
	// - Output: A B C

	x = []string{"A", "B", "C"}
	for v := range slices.Values(x) {
		fmt.Println(v)
	}

	// Collect (Seq -> slice)
	// - Collect collects values from seq into a new slice and returns it.
	// - Output: [A B C]

	x = []string{"A", "B", "C"}
	y := slices.Collect(slices.Values(x))
	fmt.Println(y)
}

//==================================================================================================
// Other Functions
//==================================================================================================

func SlicesOtherFunctions() {

	// Backward
	// - Backward returns an iterator over index-value pairs in the slice, traversing it backward
	//   with descending indices.
	// - Output: 2 C, 1 B, 0 A

	x := []string{"A", "B", "C"}
	for i, v := range slices.Backward(x) {
		fmt.Println(i, v)
	}

	// Chunk
	// - Chunk returns an iterator over consecutive sub-slices of up to n elements of s.
	// - All but the last sub-slice will have size n. All sub-slices are clipped to have no capacity
	//   beyond the length.
	// - If s is empty, the sequence is empty: there is no empty slice in the sequence. Chunk panics
	//   if n is less than 1.
	// - Output: [A B], [C D], [E]

	x = []string{"A", "B", "C", "D", "E"}
	chunk := slices.Chunk(x, 2)
	for v := range chunk {
		fmt.Println(v)
	}

	// Clip
	// - Clip removes unused capacity from the slice, returning s[:len(s):len(s)].
	// - Output: 3 (capacity is 3)

	x = make([]string, 0, 10) // (capacity is 10)
	x = append(x, "A", "B", "C")
	x = slices.Clip(x)
	fmt.Println(cap(x))

	// Grow
	// - Grow increases the slice's capacity, if necessary, to guarantee space for another n
	//   elements.
	// - Output: 5 (capacity is 5)

	x = []string{"A", "B", "C"}
	x = slices.Grow(x, 2)
	fmt.Println(cap(x))
}
