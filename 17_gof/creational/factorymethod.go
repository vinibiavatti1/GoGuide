// Factory Method
// The Factory Method pattern defines an interface for creating an object, but lets subclasses alter
// the type of objects that will be created.
// Since Go does not support OOP, this pattern looks similar to the abstract factory pattern.

package creational

// Model
// The model below is a struct that represents an Input.
// It contains a Markup field that represents the input's markup.
// This input can have different styles, such as flat or rounded.
type Input struct {
	Markup string
}

// Interface
// The interface below defines what consists of a form.
// It looks similar to the abstract factory, but it is a component itself, not a factory.
// The CreateInput method is called "Factory Method".
type Form interface {
	CreateInput() *Input
}

// Concrete Forms
// The structs below are concrete forms that implement the Form interface.
type (
	FlatForm    struct{}
	RoundedForm struct{}
)

// Form Implementations
// The functions below are implementations of a form.
// Note that the factory method is implemented with different logic for each form.
func (f *FlatForm) CreateInput() *Input {
	return &Input{Markup: "[_____]"}
}
func (f *RoundedForm) CreateInput() *Input {
	return &Input{Markup: "(_____)"}
}

// Render
// This function renders the form.
// It takes a Form interface as an argument and calls the CreateInput method to get the input markup.
func RenderForm(form Form) {
	input := form.CreateInput()
	println("Enter your name:", input.Markup)
}

// Test Factory Method
// The function below is a test for the factory method pattern.
func TestFactoryMethod() {
	flatForm := &FlatForm{}
	roundedForm := &RoundedForm{}
	RenderForm(flatForm)    // Output: Enter your name: [_____]
	RenderForm(roundedForm) // Output: Enter your name: (_____)
}
