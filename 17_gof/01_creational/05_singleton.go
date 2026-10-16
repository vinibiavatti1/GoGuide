// Singleton
// - The Singleton pattern ensures that a class has only one instance and provides a global point of
//   access to it.
// - This is useful when exactly one object is needed to coordinate actions across the system.
// - In Go, we can implement the Singleton pattern using a package-level variable and a function to
//   access it.

package creational

import "fmt"

// Singleton Struct
// - The struct below should have only an unique instance in our application.
// - To ensure that, we will use the singleton pattern to create a single instance of this struct.

type ServiceManager struct{}

// Singleton Instance
// - Declaring a package-level variable to hold the single instance of the Singleton struct.

var instance *ServiceManager

// GetInstance Function
// - Since we need only one instance of the ServiceManager struct, we will create a function that
//   returns the instance.
// - If the instance is nil, we will create a new instance of the ServiceManager struct.

func GetInstance() *ServiceManager {
	if instance == nil {
		instance = &ServiceManager{}
	}
	return instance
}

// Test Singleton
// - By comparing the returned pointer references, we can ensure that it is the same instance.
// - Output: Same address: true

func TestSingleton() {
	sm1 := GetInstance()
	sm2 := GetInstance()
	fmt.Println("Same address:", sm1 == sm2)
}
