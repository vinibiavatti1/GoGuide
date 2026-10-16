package main

import "fmt"

type List[T comparable] struct {
	Root *Node[T]
}

func NewList[T comparable]() *List[T] {
	return &List[T]{
		Root: nil,
	}
}

func (l *List[T]) Append(value T) {
	n := &Node[T]{
		Next:  nil,
		Value: value,
	}
	if l.Root == nil {
		l.Root = n
	} else {
		last := l.Last()
		last.Next = n
	}
}

func (l *List[T]) Prepend(value T) {
	n := &Node[T]{
		Next:  nil,
		Value: value,
	}
	if l.Root == nil {
		l.Root = n
	} else {
		root := l.Root
		l.Root = n
		n.Next = root
	}

}

func (l *List[T]) Last() *Node[T] {
	if l.Root == nil {
		return nil
	}
	n := l.Root
	for n.Next != nil {
		n = n.Next
	}
	return n
}

func (l *List[T]) First() *Node[T] {
	return l.Root
}

func (l *List[T]) Contains(value T) bool {
	n := l.Root
	for {
		if n == nil {
			return false
		}
		if n.Value == value {
			return true
		}
		n = n.Next
	}
}

func (l *List[T]) Len() (s int) {
	s = 0
	n := l.Root
	for {
		if n == nil {
			return
		}
		n = n.Next
		s++
	}
}

func (l *List[T]) Get(index int) *Node[T] {
	if index < 0 {
		return nil
	}
	if l.Root == nil {
		return nil
	}
	n := l.Root
	i := 0
	for {
		if i == index {
			return n
		}
		n = n.Next
		if n == nil {
			return nil
		}
		i++
	}
}

func (l *List[T]) Delete(index int) bool {
	if index < 0 {
		return false
	}
	if l.Root == nil {
		return false
	}
	if index == 0 {
		l.Root = l.Root.Next
		return true
	}
	n := l.Get(index - 1)
	if n == nil || n.Next == nil {
		return false
	}
	n.Next = n.Next.Next
	return true
}

func (l *List[T]) Pop() *Node[T] {
	i := l.Len() - 1
	n := l.Get(i - 1)
	if n == nil {
		return nil
	}
	r := n.Next
	n.Next = nil
	return r
}

func (l *List[T]) Empty() bool {
	if l.Root == nil {
		return true
	}
	return false
}

func (l *List[T]) Clear() {
	l.Root = nil
}

func (l *List[T]) Print() {
	if l.Root == nil {
		fmt.Println("<empty>")
		return
	}
	n := l.Root
	i := 0
	for {
		fmt.Println(i, ":", n.Value)
		i++
		n = n.Next
		if n == nil {
			return
		}
	}
}

type Node[T any] struct {
	Next  *Node[T]
	Value T
}

func main() {
	list := NewList[string]()
	list.Append("A")
	list.Append("B")
	list.Append("C")
	list.Prepend("D")
	list.Prepend("E")
	list.Print()
	l := list.Len()
	fmt.Println("Len:", l)
	n := list.Get(3)
	fmt.Println("Node:", n)
	p := list.Pop()
	fmt.Println("Node:", p)
	list.Print()
}
