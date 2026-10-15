// Facade
// The Facade pattern provides a simplified interface to a complex subsystem.
// It defines a higher-level interface that makes the subsystem easier to use.

package structural

// Complex Subsystem
// We will declare a complex subsystem (couple of interfaces) that will be used by the Facade.
// These interfaces are related to a common operation.
type DBConnector interface {
	Connect()
}
type MessageLogger interface {
	Config()
}
type MessageListener interface {
	Listen()
}

// Facade
// The struct below represents a Facade.
// It will provide a simplified interface to the complex subsystem.
// The Facade will use the complex subsystem to perform a common operation.
type MessageProcessorFacade struct {
	dbConnector     DBConnector
	messageLogger   MessageLogger
	messageListener MessageListener
}

// Facade Implementation
// Now, we will implement a common operation that will be used by the Facade.
// This avoids the consumer from needing to know the details of the complex subsystem.
func (f *MessageProcessorFacade) Bootstrap() {
	f.dbConnector.Connect()
	f.messageLogger.Config()
	f.messageListener.Listen()
}

// Test Facade
// The function below will test the Facade.
func TestFacade() {
	mp := MessageProcessorFacade{}
	mp.Bootstrap()
}
