// Template Method
// The Template Method is a behaviora pattern that defines the skeleton of an algorithm in a method, deferring
// some steps to subclasses.
// It lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.

package behavioral

import "fmt"

// Interface
// The interface defines the skeleton of an algorithm, allowing subclasses to override specific steps without
// changing the overall structure.
type EnemyAI interface {
	Turn()            // Template method
	CollectResource() // Step 1
	BuildStructure()  // Step 2
	BuildUnit()       // Step 3
}

// Base
// The base struct implements the template method and provides default implementations for some steps
// of the algorithm.
type BaseAI struct{}

// Base Implementation
// Some steps of the algorithm are implemented in the base class, while others are left to be overridden by subclasses.
// The template method defines the skeleton of the algorithm, calling the steps in a specific order.
func (a *BaseAI) Turn() {
	a.CollectResource()
	a.BuildStructure()
	a.BuildUnit()
}
func (a *BaseAI) BuildStructure() {
	fmt.Println("Structure Built!")
}
func (a *BaseAI) BuildUnit() {
	fmt.Println("Unit Built!")
}
func (a *BaseAI) CollectResource() {
	fmt.Println("Gold Collected!")
}

// Children
// The subclasses override specific steps of the algorithm to provide their own implementations.
// They can also call the base class's implementation if needed.
type (
	OrcsAI   struct{ BaseAI }
	HumansAI struct{ BaseAI }
)

// Overriding Methods (Orcs)
// The subclasses override the methods to provide their own implementations of the steps in the algorithm.
// Note that some steps are not overridden, so the base class's implementation will be used.
func (a *OrcsAI) BuildStructure() {
	fmt.Println("Orc Structure Built!")
}
func (a *OrcsAI) BuildUnit() {
	fmt.Println("Orc Unit Built!")
}

// Overriding Methods (Humans)
// Now, we will override the methods in the HumansAI struct to provide different implementations.
func (a *HumansAI) BuildStructure() {
	fmt.Println("Human Structure Built!")
}
func (a *HumansAI) CollectResource() {
	fmt.Println("Food Collected!")
}

// Test Template Method
// The function demonstrates how the template method works with different subclasses.
// It shows how the subclasses can provide their own implementations of the steps in the algorithm while
// still using the base class's implementation for other steps.
func TestTemplateMethod() {
	oai := &OrcsAI{}
	hai := &HumansAI{}

	oai.Turn()
	// Output: Orcs Turn:
	// Gold Collected!        (Base)
	// Orc Structure Built!   (Overrided)
	// Orc Unit Built!        (Overrided)

	hai.Turn()
	// Output: Humans Turn:
	// Food Collected!        (Overrided)
	// Human Structure Built! (Overrided)
	// Unit Built!            (Base)
}
