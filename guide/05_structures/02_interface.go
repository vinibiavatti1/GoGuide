// Interface
// Interfaces are used to define a set of methods that a type must implement.
// They provide a way to specify the behavior of an object: if something can do this, then it can be used here.
// There is also Type Interface, which is a type that specifies a set of types that a value can have.

package datatypes

import "fmt"

// Interface
// An interface is a type that specifies a set of methods that a concrete type must implement.
type Printer interface {
	Print()
}

// Composition
// Interfaces can be composed of other interfaces.
// This means that an interface can include other interfaces as part of its definition.
type OtherPrinter interface {
	Printer
	OtherPrint()
}

// Nil Interface
// Interfaces can be nil, which means that they do not hold any value.
var _ Printer = nil

// Defining a Struct
// Below we will define a struct that to implement the Printer interface.
type Data struct {
	Key   string
	Value string
}

// Implementing Interface
// A type implements an interface by implementing its methods.
// The interface is implicitly implemented by any type that implements all the methods.
// We don't need to explicitly declare that a type implements an interface.
func (d *Data) Print() {
	fmt.Printf("%s: %s", d.Key, d.Value)
}

// Using Interface Implementation
// Now we can use the interface to call the method on any type that implements it.
func CheckInterfaceImplementation() {

	// Creating an Instance
	// We will create an instance of the Data struct.
	data := &Data{Key: "A", Value: "1"}

	// Calling the Interface Method
	// Since Data implements the Printer interface, we can call the Print method on it.
	data.Print() // Output: A: 1

	// Ensuring Interface Implementation
	// To ensure that the Data struct implements the Printer interface, we can use a type assertion.
	// This will cause a compile-time error if Data does not implement Printer.
	// This is a common pattern in Go to ensure that a type implements an interface.
	var _ Printer = (*Data)(nil) // This will not cause an error if Data implements Printer.
}

// Local Interfaces
// We can define local interfaces inside functions.
func LocalInterfaces() {

	// Creating an Instance
	// We will create an instance of the Data struct.
	data := &Data{Key: "A", Value: "1"}
	fmt.Println("Data:", data) // Output: Data: &{A 1}

	// Defining an Interface
	// We can define an interface inside a function.
	// This interface will only be available inside the function.
	type otherPrinter interface {
		Print()
	}

	// Ensuring Interface Implementation
	// Knowing that the "otherPrinter" interface has the same signature as the "Printer" interface, we can
	// check if the Data struct implements it.
	var _ otherPrinter = (*Data)(nil) // This will not cause an error if Data implements otherPrinter.
}
