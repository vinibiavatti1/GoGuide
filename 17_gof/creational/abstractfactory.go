// Abstract Factory
// Abstract Factory is a creational design pattern that lets you produce families of related objects without
// specifying their concrete classes.
// Since Go does not support OOP, this pattern looks similar to the factory method pattern.

package creational

import "fmt"

// Model
// The model below is a struct that represents a button.
// It contains a Markup field that represents the button's markup.
// This button can have different styles, such as flat or rounded.
type Button struct {
	Markup string
}

// Factory
// The interface below defines the factory for creating buttons.
// It contains a method CreateButton that returns a pointer to a Button.
type ButtonFactory interface {
	CreateButton() *Button
}

// Concrete Factories
// The structs below are concrete factories that implement the ButtonFactory interface.
// The FlatButtonFactory creates flat buttons, while the RoundedButtonFactory creates rounded buttons.
type (
	FlatButtonFactory    struct{}
	RoundedButtonFactory struct{}
)

// Factory Implementations
// The functions below are implementations for the factories.
// The FlatButtonFactory creates a button with flat markup.
// The RoundedButtonFactory creates a button with rounded markup.
func (d *FlatButtonFactory) CreateButton() *Button {
	return &Button{Markup: "[Button]"}
}
func (d *RoundedButtonFactory) CreateButton() *Button {
	return &Button{Markup: "(Button)"}
}

// Test Factory
// The function below is a test for the abstract factory pattern.
func TestFactory() {
	flatFactory := &FlatButtonFactory{}
	roundFactory := &RoundedButtonFactory{}
	b1 := flatFactory.CreateButton()
	b2 := roundFactory.CreateButton()
	fmt.Println("b1:", b1.Markup, "b2:", b2.Markup) // Output: b1: [Button] b2: (Button)
}
