// Strategy
// The Strategy pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable.
// Strategy lets the algorithm vary independently from clients that use it.
// It is a behavioral design pattern that enables selecting an algorithm's behavior at runtime.

package behavioral

import "fmt"

// Strategy
// The Strategy interface declares a method for executing a strategy.
// The Context uses this interface to call the algorithm defined by Concrete Strategies.
type Strategy interface {
	Check(number int) bool
}

// Context
// The Context defines the interface of interest to clients.
// It maintains a reference to one of the Strategy objects and delegates it executing the algorithm.
// The Context does not know the concrete class of a strategy. It should work with all strategies via the
// Strategy interface.
type EvenChecker struct {
	Strategy Strategy
}

// Context Implementation
// Here, the Context delegates some work to the Strategy object instead of implementing multiple versions
// of the algorithm on its own.
func (d *EvenChecker) Check(number int) bool {
	return d.Strategy.Check(number)
}

// Concrete Strategies
// Concrete Strategies implement the algorithm while following the base Strategy interface.
// The interface makes them interchangeable in the Context.
// The client should be able to use any Concrete Strategy without knowing the details of the implementation.
type (
	ArithmeticStrategy struct{}
	BinaryStrategy     struct{}
)

// Concrete Strategy Implementation
// The concrete strategies will have different implementations of the Check method.
// The client can choose which strategy to use at runtime.
func (s *ArithmeticStrategy) Check(number int) bool {
	return number%2 == 0
}
func (s *BinaryStrategy) Check(number int) bool {
	return number&1 == 0
}

// Test Strategy
// The function demonstrates how the Context can be configured with different strategies at runtime.
func TestStrategy() {
	as := &ArithmeticStrategy{}
	bs := &BinaryStrategy{}
	d := &EvenChecker{Strategy: as}
	fmt.Println("Using Arithmetic Strategy:", d.Check(4)) // Output: Using Arithmetic Strategy: true
	d.Strategy = bs
	fmt.Println("Using Binary Strategy:", d.Check(4)) // Output: Using Binary Strategy: true
}
