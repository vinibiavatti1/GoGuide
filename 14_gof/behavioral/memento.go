// Memento
// The Memento pattern is a behavioral design pattern that allows an object to capture its internal state
// and save it externally, so that it can be restored later without violating encapsulation.
// The Memento pattern is often used to implement undo functionality in applications, allowing users to revert
// to a previous state of an object without exposing its internal structure.

package behavioral

import "fmt"

// Model
// The Model represents the object whose state we want to save and restore.
// In this case, it is a simple Document object with a title and content.
type Document struct {
	Content string
	Title   string
}

// Memento
// The Memento is the object that stores the internal state of the Document object.
// It is used to capture the state of the Document at a specific point in time.
type DocumentSnapshot struct {
	Content string
	Title   string
}

// Caretaker
// The Caretaker is responsible for managing the Memento objects and providing the ability to save and restore
// the state of the Document.
// In this case, the Editor acts as the Caretaker, holding a reference to the Document and the Memento.
type Editor struct {
	Document *Document
	Snapshot *DocumentSnapshot
}

// Implementation
// The Editor class provides methods to save the current state of the Document as a Memento and to restore it later.
// The SaveSnapshot method creates a new Memento object and saves the current state of the Document.
func (e *Editor) SaveSnapshot() {
	e.Snapshot = &DocumentSnapshot{
		Content: e.Document.Content,
		Title:   e.Document.Title,
	}
}
func (e *Editor) Undo() {
	if e.Snapshot != nil {
		e.Document.Content = e.Snapshot.Content
		e.Document.Title = e.Snapshot.Title
	}
}

// Test Memento
// The TestMemento function demonstrates the Memento pattern by creating a simple Editor object with a Document.
func TestMemento() {
	e := &Editor{
		Document: &Document{Title: "New Document", Content: "..."},
		Snapshot: nil,
	}
	e.SaveSnapshot()
	e.Document.Title = "Hello World!"
	e.Document.Content = "Lorem ipsum dolor"

	fmt.Println(e.Document.Title, e.Document.Content) // Output: Hello World! Lorem ipsum dolor
	e.Undo()                                          // Restore the previous state
	fmt.Println(e.Document.Title, e.Document.Content) // Output: New Document ...
}
