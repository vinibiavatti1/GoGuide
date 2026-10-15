// Decorator
// The Decorator pattern is a structural pattern that allows behavior to be added to individual objects,
// either statically or dynamically, without affecting the behavior of other objects from the same class.
// It is a flexible alternative to subclassing for extending functionality.
// In Go, the Decorator pattern can be implemented using interfaces and struct embedding.

package structural

import "fmt"

// Protocol
// The interface below defines a default protocol for email handlers.
// The concrete implementation of the protocol will be set on the main model and on the decorators.
type EmailHandler interface {
	Append(content string)
	Send()
}

// Model
// The model is the main object that will be decorated.
// It implements the EmailHandler interface and contains the core functionality.
type Email struct {
	content string
}

// Model Implementation
// Here, we implement the EmailHandler interface for the Email struct.
// The Append method adds content to the email, and the Send method prints the email content.
func (e *Email) Append(content string) {
	e.content += content
}
func (e *Email) Send() {
	fmt.Println(e.content)
}

// Decorators
// The decorators are structs that also implement the EmailHandler interface.
// They contain a reference to an EmailHandler and add additional functionality to the Send method.
type (
	SignedEmailDecorator   struct{ Email EmailHandler }
	AttachedEmailDecorator struct{ Email EmailHandler }
)

// Decorator Implementation
// The decorators implement the EmailHandler interface and add their own behavior to the Send method.
func (e *SignedEmailDecorator) Append(content string) {
	e.Email.Append(content)
}
func (e *SignedEmailDecorator) Send() {
	e.Email.Append("\nSigned by: John Duo")
	e.Email.Send()
}

// Decorator Implementation
// Here, other decorator is implemented.
// It will add its own behavior to the Send method, including the attachment.
func (e *AttachedEmailDecorator) Append(content string) {
	e.Email.Append(content)
}
func (e *AttachedEmailDecorator) Send() {
	e.Email.Append("\nAttach: File.txt")
	e.Email.Send()
}

// Test Decorator
// The test function demonstrates how to use the decorator pattern.
func TestDecorator() {
	var e EmailHandler = &Email{content: "Hello World!"}
	e = &AttachedEmailDecorator{Email: e} // Decorate with attachment
	e = &SignedEmailDecorator{Email: e}   // Decorate with signature
	e.Send()
	// Output:
	//
	// Hello World!
	// Signed by: John Duo
	// Attach: File.txt
}
