// Mediator
// Mediator is a behavioral design pattern that allows objects to communicate with each other without knowing about each other.
// It defines an object that encapsulates how a set of objects interact.
// This pattern is useful when you want to reduce the complexity of communication between multiple objects.

package behavioral

import "fmt"

// Interface
// The interface below will be used to demonstrate the Mediator pattern.
type Checkbox interface {
	OnCheck()
	OnUncheck()
}

// Concrete Implementation
// This is a concrete implementation of the Checkbox interface.
type ConsentCheckbox struct {
	Parent *Dialog
}

// Method Implementations
// Note that the logic of the methods is not in the concrete implementation, but in the parent object.
// This is the key to the Mediator pattern, since the dialog will be the mediator.
func (c *ConsentCheckbox) OnCheck() {
	c.Parent.Notify("enableSubmit")
}
func (c *ConsentCheckbox) OnUncheck() {
	c.Parent.Notify("disableSubmit")
}

// Mediator
// The mediator is the object that encapsulates how a set of objects interact.
// In this case, the dialog is the mediator.
type Dialog struct {
	SubmitButtonEnabled bool
}

// Mediator Notifier
// The notify method is the method that will be called by all the objects that want to communicate with the mediator.
// In this case, the checkbox will call the notify method of the dialog when it is checked or unchecked.
// The dialog will then enable or disable the submit button based on the state of the checkbox.
func (d *Dialog) Notify(event string) {
	switch event {
	case "enableSubmit":
		d.SubmitButtonEnabled = true
	case "disableSubmit":
		d.SubmitButtonEnabled = false
	}
}

// Test Mediator
// The test function will create a dialog and a checkbox, and then check the checkbox.
func TestMediator() {
	dialog := &Dialog{}
	checkbox := &ConsentCheckbox{Parent: dialog}
	fmt.Println(dialog.SubmitButtonEnabled) // Output: false
	checkbox.OnCheck()
	fmt.Println(dialog.SubmitButtonEnabled) // Output: true
}
