// Builder
// The Builder pattern is a creational design pattern that allows for the step-by-step construction of complex objects.
// It separates the construction of a complex object from its representation,
// allowing the same construction process to create different representations.

package creational

import "fmt"

// Model
// The model below will be used to demonstrate the Builder pattern.
type Person struct {
	ID   int
	Name string
}

// Director
// The director is responsible for managing the construction process.
// It uses a builder to construct the object step by step.
type PersonDirector struct {
	Builder PersonBuilder
}

// Director Function
// The function below is an implementation for the director.
func (d *PersonDirector) Build(id int, name string) *Person {
	pb := d.Builder
	pb.WithID(id)
	pb.WithName(name)
	return pb.Build()
}

// Builder Interface
// The builder interface defines the methods for constructing the object.
// It allows for different implementations of the builder to create different representations of the object.
// The builder interface is used by the director to construct the object step by step.
type PersonBuilder interface {
	WithID(id int)
	WithName(name string)
	Build() *Person
}

// Builder
// The default builder is a concrete implementation of the builder interface.
// It provides the methods implementation for constructing the object step by step.
type DefaultPersonBuilder struct {
	Person *Person
}

// Builder Functions
// The functions below are implementations for the builder.
func (b *DefaultPersonBuilder) WithID(id int) {
	b.Person.ID = id
}
func (b *DefaultPersonBuilder) WithName(name string) {
	b.Person.Name = name
}
func (b *DefaultPersonBuilder) Build() *Person {
	return b.Person
}

// Test Builder
// The test function below demonstrates the usage of the builder pattern.
// It creates a new person using the builder and prints the result.
func TestBuilder() {
	pb := &DefaultPersonBuilder{Person: &Person{}}
	pd := &PersonDirector{Builder: pb}
	p := pd.Build(1, "John Doe")
	fmt.Println(p) // Output: &{1 John Doe}
}
