// Iterators
// Iterators provide a way to access the elements of an aggregate object sequentially without exposing
// its underlying representation.
// The "iter" package provides an implementation of the iterator pattern.
// In Go, Iterators are implemented with a function that passes successive elements of a sequence to a callback function,
// conventionally named yield. This pattern is accepted by range loops.
// The function stops either when the sequence is finished or when yield returns false,
// indicating to stop the iteration early.
// See: https://go.dev/blog/range-functions

package gopatterns

import (
	"iter"
)

// The Yield Function
// As of Go 1.23 it now supports ranging over functions that take a single argument.
// The single argument must itself be a function that takes zero to two arguments and returns a bool.
// By convention, we call it the yield function.
// This function allows us to iterate over a sequence of values, witout the need of exposing
// its underlying representation.
// When we speak of an iterator in Go, we mean a function with one of these three types:
var _ = `
func(yield func() bool)
func(yield func(V) bool)
func(yield func(K, V) bool)
`

// Sequences (Seq and Seq2)
// To make iterators easier to use, the new standard library package iter defines two types: Seq and Seq2.
// These are names for the iterator function types, the types that can be used with the for/range statement.
// The difference between Seq and Seq2 is just that Seq2 is a sequence of pairs, such as a key and a value from a map.
var _ = `
type (
	Seq[V any]     func(yield func(V) bool)
	Seq2[K, V any] func(yield func(K, V) bool)
)
`

// Creating a Custom Slice Type
// For the iterator examples, we will create a custom type called List.
// This type will be based on a slice of any type.
// After the declaration, we will implement the iterator function for this type, which
// will allow the type to be used with the for/range statement.
type List[T any] []T

// All Function (Convention for iterators)
// The default name for a function that returns a Seq or Seq2 is All.
// This is a convention, not a requirement.
// The name All is used to indicate that the function returns all elements of the sequence.
// The function needs to return a Seq or Seq2 type, to be used with the for/range statement.
// The function below is the default implementation of a iterator.
func (l *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range *l {
			if !yield(v) {
				break
			}
		}
	}
}

// Iterating Over a List
// Since the List type implements the iterator function, we can use it with the for/range statement.
// The example below shows how to iterate over a List of integers.
func IterateOverList() {

	// Creating a List
	// The List type is based on a slice of integers.
	x := List[int]{1, 2, 3, 4, 5}

	// Iterating Over the List
	// The for/range statement is used to iterate over the List.
	for v := range x.All() {
		println(v) // Output: 1 2 3 4 5
	}
}

// Creating a Custom Map Type
// We can also create a custom type based on a map to explore the Seq2 type.
// The Seq2 type is used to iterate over a sequence of pairs, such as a key and a value from a map.
type Map[K comparable, V any] map[K]V

// All Function (Convention for iterators)
// Now, we will implement the iterator function for the Map type.
func (m *Map[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range *m {
			if !yield(k, v) {
				return
			}
		}
	}
}

// Iterating Over a Map
// The example below shows how to iterate over a Map of integers and strings.
func IterateOverMap() {

	// Creating a Map
	// The Map type is based on a map of integers and strings.
	x := Map[int, string]{1: "A", 2: "B", 3: "C"}

	// Iterating Over the Map
	// The for/range statement is used to iterate over the Map.
	for k, v := range x.All() {
		println(k, v) // Output: 1 A 2 B 3 C
	}
}
