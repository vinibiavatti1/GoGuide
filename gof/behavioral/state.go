// State
// The State pattern allows an object to change its behavior when its internal state changes.
// The object will appear to change its class.
// This pattern is useful when you want to avoid using large conditional statements to manage state transitions.
// It allows you to encapsulate state-specific behavior in separate classes, making the code more maintainable
// and easier to understand.

package behavioral

import "fmt"

// State
// The State interface defines the methods that will be implemented by the concrete states.
// In this case, it is the Edit, Publish, and Unpublish methods that will be called when the post changes state.
type PostState interface {
	Edit(content string)
	Publish()
	Unpublish()
}

// Context
// The Post struct is the context that will change its state based on the current state.
// It contains a reference to the current state and the content of the post.
type Post struct {
	State   PostState
	Content string
}

// Constructor
// The NewPost function creates a new post and sets its initial state to DraftState.
func NewPost() *Post {
	post := &Post{}
	post.State = &DraftState{Post: post}
	return post
}

// Context Implementation
// The Post struct has methods to change its state and edit its content.
// The Edit method calls the Edit method of the current state, and the Publish and Unpublish methods
// call the corresponding methods of the current state.
func (p *Post) Edit(content string) {
	p.State.Edit(content)
}
func (p *Post) Publish() {
	p.State.Publish()
}
func (p *Post) Unpublish() {
	p.State.Unpublish()
}

// Concrete States
// The DraftState and PublishedState structs are concrete implementations of the PostState interface.
// They define the behavior of the post when it is in the draft or published state, respectively.
type (
	DraftState     struct{ Post *Post }
	PublishedState struct{ Post *Post }
)

// Draft State Implementation
// The DraftState struct implements the PostState interface and defines the behavior of the post when it
// is in the draft state.
func (p *DraftState) Edit(content string) {
	p.Post.Content = content
	fmt.Println("Post edited")
}
func (p *DraftState) Publish() {
	p.Post.State = &PublishedState{Post: p.Post}
	fmt.Println("Post published")
}
func (p *DraftState) Unpublish() {
	fmt.Println("Cannot unpublish a draft post")
}

// Published State Implementation
// The PublishedState struct implements the PostState interface and defines the behavior of the post when it
// is in the published state.
func (p *PublishedState) Edit(content string) {
	fmt.Println("Cannot edit a published post")
}
func (p *PublishedState) Publish() {
	fmt.Println("Post is already published")
}
func (p *PublishedState) Unpublish() {
	p.Post.State = &DraftState{Post: p.Post}
	fmt.Println("Post unpublished")
}

// Test State
// The test function will create a post and change its state from draft to published and back to draft.
func TestState() {
	post := NewPost()                          // Output:
	post.Edit("Hello, World!")                 // Post edited
	post.Publish()                             // Post published
	post.Edit("Hello, Universe!")              // Cannot edit a published post
	post.Unpublish()                           // Post unpublished.
	post.Edit("Hello, Galaxy!")                // Post edited
	post.Publish()                             // Change state to published
	post.Publish()                             // Post is already published
	fmt.Println("Post content:", post.Content) // Post content: Hello, Galaxy!
}
