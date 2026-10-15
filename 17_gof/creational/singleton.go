// Singleton
// The Singleton pattern ensures that a class has only one instance and provides a global point of access to it.
// This is useful when exactly one object is needed to coordinate actions across the system.
// In Go, we can implement the Singleton pattern using a package-level variable and a function to access it.

package creational

import "fmt"

// Singleton Struct
// The struct below should have only an unique instance in our application.
// To ensure that, we will use the singleton pattern to create a single instance of this struct.
type ServiceManager struct {
	ID int
}

// Singleton Instance
// Declaring a package-level variable to hold the single instance of the Singleton struct.
var instance *ServiceManager

// Singleton Constructor
// Since we need only one instance of the ServiceManager struct, we will create a constructor function
// that returns the instance.
// If the instance is nil, we will create a new instance of the ServiceManager struct.
func NewServiceManager() *ServiceManager {
	if instance == nil {
		instance = &ServiceManager{ID: 1}
	}
	return instance
}

// Test Singleton
// We can check if the instance is the same by comparing the IDs of the instances returned by GetInstance.
func TestSingleton() {

	// Getting the Singleton instance
	// We will call the NewSingleton function twice to validate if both instances are the same.
	sm1 := NewServiceManager()
	sm2 := NewServiceManager()

	// Check
	// We can see that both instances have the same ID, which means they are indeed the same instance.
	fmt.Println("Singleton Instance IDs:", sm1.ID, sm2.ID) // Output: Singleton Instance IDs: 1 1
}
