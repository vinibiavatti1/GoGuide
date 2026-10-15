// Iterator
// The Iterator pattern is a behavioral design pattern that provides a way to access the elements of an
// aggregate object sequentially without exposing its underlying representation.
// It consists of two main components: the Iterator and the Iterable (or Aggregate).
// The Iterator is responsible for iterating over the elements, while the Iterable provides a way to create an Iterator.

package behavioral

import "fmt"

// Iterator
// The Iterator interface defines the methods that an iterator must implement.
// It typically includes methods for getting the next element and checking if there are more elements to iterate over.
type Iterator[T any] interface {
	Next() T
	HasMore() bool
}

// Concrete Iterator
// The ConcreteIterator is a specific implementation of the Iterator interface.
// It holds a reference to the aggregate object (in this case, a List) and keeps track of the current position
// in the iteration.
type ListIterator[T any] struct {
	List  *List[T]
	index int
}

// Implementation
// The ListIterator implements the Iterator interface and provides methods to iterate over the elements of the List.
// The Next method returns the next element in the iteration, and the HasMore method checks if there are more elements
// to iterate over.
func (i *ListIterator[T]) Next() (res T) {
	res = i.List.Data[i.index]
	i.index++
	return
}
func (i *ListIterator[T]) HasMore() bool {
	return i.index < len(i.List.Data)
}

// Iterable
// The Iterable is the collection or aggregate object that provides a way to create an iterator.
// In this case, it is a List that holds a slice of elements and provides a method to create an iterator.
type List[T any] struct {
	Data []T
}

// Implementation
// The List struct holds a slice of elements and provides a method to create an iterator.
// The Iterator method returns a new ListIterator that can be used to iterate over the elements of the List.
func (l *List[T]) Iterator() Iterator[T] {
	return &ListIterator[T]{
		List:  l,
		index: 0,
	}
}

// Test Iterator
// The TestIterator function demonstrates the Iterator pattern by creating a simple List object and iterating
// over its elements.
func TestIterator() {
	list := &List[string]{}
	list.Data = append(list.Data, "A", "B", "C")
	iter := list.Iterator()
	for iter.HasMore() {
		fmt.Println(iter.Next()) // Output: A B C
	}
}
