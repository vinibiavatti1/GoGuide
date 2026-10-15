// Observer
// The Observer pattern defines a one-to-many dependency between objects so that when one object changes
// state, all its dependents are notified and updated automatically.
// This pattern is useful when you want to maintain a consistent state across multiple objects without
// tightly coupling them.

package behavioral

import "fmt"

// Observer
// The Observer interface defines the method that will be called when the subject changes state.
// In this case, it is the Notify method that will be called when the MessageProcessor processes a message.
type Observer interface {
	Notify(msg *Message)
}

// Publisher
// The Publisher is the subject that maintains a list of subscribers and notifies them when its state changes.
// In this case, it is the MessageProcessor that processes messages and notifies its subscribers.
type MessageProcessor struct {
	Subscribers []Observer
}

// Publisher Implementation
// The Subscribe method adds a new subscriber to the list of subscribers.
// The Process method processes a message and notifies all subscribers.
func (m *MessageProcessor) Subscribe(s Observer) {
	m.Subscribers = append(m.Subscribers, s)
}
func (m *MessageProcessor) Process(msg *Message) {
	fmt.Println("[MessageProcessor] Message Processed:", msg.Content)
	for _, v := range m.Subscribers {
		v.Notify(msg)
	}
}

// Context
// The Message struct represents the context that will be processed by the MessageProcessor and sent to the subscribers.
// In this case, it is a simple struct with a Content field.
type Message struct {
	Content string
}

// Concrete Observers
// The LoggingService and DataService are concrete implementations of the Observer interface.
// They define the Notify method that will be called when the MessageProcessor processes a message.
type (
	LoggingService struct{}
	DataService    struct{}
)

// Concrete Observers Implementation
// The Notify method is called when the MessageProcessor processes a message.
// In this case, it simply prints the message content to the console.
func (s *LoggingService) Notify(msg *Message) {
	// Process Message...
	fmt.Println("[LoggingService] Received:", msg.Content)
}
func (s *DataService) Notify(msg *Message) {
	// Process Message...
	fmt.Println("[DataService] Received:", msg.Content)
}

// Test Observer
// The TestObserver function demonstrates the usage of the Observer pattern.
func TestObserver() {
	mp := &MessageProcessor{}
	ls := &LoggingService{}
	ds := &DataService{}
	mp.Subscribe(ls) // Subscribe LoggingService
	mp.Subscribe(ds) // Subscribe DataService
	m := &Message{Content: "Hello World!"}
	mp.Process(m) // Process Message and Notify Subscribers
	// Output:
	// [MessageProcessor] Message Processed: Hello World!
	// [LoggingService] Received: Hello World!
	// [DataService] Received: Hello World!
}
