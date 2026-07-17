// Bridge
// The Bridge pattern is a structural design pattern that decouples an abstraction from its implementation,
// allowing the two to vary independently. It is used to separate the abstraction from the implementation,
// enabling you to change the implementation without affecting the abstraction and vice versa.

package structural

import (
	"fmt"
)

// Interface
// The interface defines the common operations for the concrete implementations.
type Writer interface {
	Write(content string)
}

// Struct
// This is the struct that will implement the interface.
type File struct {
	Content string
}

// Concrete implementation
// The concrete implementation of the interface.
func (f *File) Write(content string) {
	f.Content += content
}

// Bridge
// The bridge struct contains a reference to the interface.
// This is the abstraction that will use the implementation.
type HTMLWriter struct {
	Writer Writer
}

// Bridge Implementation
// The bridge implementation uses the interface to write content.
func (w *HTMLWriter) Paragraph(content string) {
	w.Writer.Write("<p>" + content + "</p>")
}

// Test Bridge
// The test function creates a concrete implementation and a bridge.
func TestBridge() {
	impl := &File{}                    // Implementation
	bridge := HTMLWriter{Writer: impl} // Abstraction (Bridge)
	bridge.Paragraph("Hello World!")   // ...
	fmt.Println(impl.Content)          // Output: <p>Hello World!</p>
}
