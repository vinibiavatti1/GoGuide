// Command
// The Command pattern is a behavioral design pattern that turns a request into a stand-alone object.
// The pattern encapsulates a request as an object, thereby allowing for parameterization of clients
// with queues, requests, and operations.
// It also provides support for undoable operations.

package behavioral

import "fmt"

// Command
// The Command interface declares a method for executing a command.
type Command interface {
	Execute()
}

// Receiver
// The Receiver is the object that knows how to perform the operations associated with carrying out a request.
// In this case, it is the application that performs the operations.
// The Receiver can have a history of commands, to allow for undo operations or to keep track of executed commands.
type App struct {
	History []Command
}

// Invoker
// The Invoker is responsible for holding and executing the command.
type Button struct {
	Command Command
}

// Invoker Implementation
// Here, we will define a method to run the command.
// The Button acts as the invoker that calls the command's execute method.
func (b *Button) Click() {
	b.Command.Execute()
}

// Concrete Commands
// The Concrete Command classes implement the Command interface and define the binding between a Receiver and an action.
// In this case, we have two commands: OpenCommand and CloseCommand.
type (
	OpenCommand  struct{ App *App }
	CloseCommand struct{ App *App }
)

// Command Implementation
// The Execute method is where the command is executed.
// In this case, we will just print the command name to the console.
func (c *OpenCommand) Execute() {
	fmt.Println("Open")
	c.App.History = append(c.App.History, c)
}
func (c *CloseCommand) Execute() {
	fmt.Println("Close")
	c.App.History = append(c.App.History, c)
}

// Test Command
// The TestCommand function demonstrates the Command pattern by creating a simple application with
// buttons that execute commands.
func TestCommand() {
	app := &App{History: []Command{}}

	oc := &OpenCommand{App: app}
	cc := &CloseCommand{App: app}

	b1 := &Button{Command: oc} // OpenCommand
	b2 := &Button{Command: oc} // Reusing the OpenCommand command
	b3 := &Button{Command: cc} // CloseCommand

	b1.Click() // Output: Open
	b2.Click() // Output: Open
	b3.Click() // Output: Close

	// Note that the history of commands is stored in the App struct.
	for _, cmd := range app.History {
		cmd.Execute() // Output: Open, Open, Close
	}
}
