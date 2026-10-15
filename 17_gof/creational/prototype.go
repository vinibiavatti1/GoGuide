// Prototype
// The Prototype pattern is a creational design pattern that allows you to create new objects by
// copying an existing object, known as the prototype.
// This pattern is useful when the cost of creating a new object is more expensive than copying an
// existing one.
// Usually, to implement the Prototype pattern, we need to implement a clone method.

package creational

import "fmt"

// Struct
// We will declare a struct that can be cloned.
type Payment struct {
	Amount float64
	Tax    float64
}

// Clone
// The method below is an implementation of the Clone method into the Payment struct.
// It will create a new object by copying the existing one.
func (p *Payment) Clone() *Payment {
	return &Payment{
		Amount: p.Amount,
		Tax:    p.Tax,
	}
}

// Test Prototype
// The function below will test the Prototype pattern.
func TestPrototype() {
	p1 := &Payment{
		Amount: 100.0,
		Tax:    10.0,
	}
	p2 := p1.Clone()
	p2.Amount = 200.0
	fmt.Println("p1:", p1, "p2:", p2) // Output: p1: &{100 10} p2: &{200 10}
}
