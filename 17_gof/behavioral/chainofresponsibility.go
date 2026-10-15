// Chain of Responsibility
// The Chain of Responsibility pattern is a behavioral design pattern that allows an object to pass a request
// along a chain of potential handlers until one of them handles the request.
// This pattern decouples the sender and receiver of a request, allowing multiple objects to handle the request
// without the sender needing to know which object will handle it.

package behavioral

import "fmt"

// Protocol
// The interface below defines a handler protocol.
// Each handler will implement this interface and will be responsible for processing the request.
type Handler interface {
	SetNext(handler Handler)
	Handle(request *Request)
}

// Base Handler
// The base handler will be used by all concrete handlers to facilitate the implementation of the chain.
type BaseHandler struct {
	Next Handler
}

// Base Handler Implementation
// The base handler will implemente methods that will be used by all concrete handlers.
func (h *BaseHandler) SetNext(handler Handler) {
	h.Next = handler
}
func (h *BaseHandler) CallNext(request *Request) {
	if h.Next != nil {
		h.Next.Handle(request)
	}
}

// Handlers
// The handlers below are concrete handlers that will process the request.
// Each handler will call the next handler if it is set, allowing the request to be passed along the chain.
type (
	ContentHandler struct{ BaseHandler }
	HeaderHandler  struct{ BaseHandler }
)

// Handlers Implementation
// The handlers implement the Handle method, which processes the request.
// Each handler will modify the request in some way and then call the next handler in the chain.
func (h *ContentHandler) Handle(request *Request) {
	request.Content += " [Signed]"
	h.CallNext(request)
}
func (h *HeaderHandler) Handle(request *Request) {
	request.Header = "Signed=True"
	h.CallNext(request)
}

// Request
// Base model to be used by the handlers.
type Request struct {
	Header  string
	Content string
}

// Test Chain of Responsibility
// The test function below demonstrates how to use the Chain of Responsibility pattern.
// It creates a chain of handlers and processes a request through the chain.
func TestChainOfResponsibility() {
	r := &Request{
		Header:  "",
		Content: "Hello World!",
	}
	h1 := &ContentHandler{}
	h2 := &HeaderHandler{}
	h1.SetNext(h2)
	h1.Handle(r)
	fmt.Println("Request:", r) // Output: Request: &{Signed=True Hello World! [Signed]}
}
