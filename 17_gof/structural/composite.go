// Composite
// Composite is a structural design pattern that lets you compose objects into tree structures to
// represent part-whole hierarchies.
// Composite lets clients treat individual objects and compositions of objects uniformly.
// In other words, it allows you to create a tree structure where each node can be either a leaf or a composite.

package structural

import "fmt"

// Interface
// The interface defines the common operations for both leaf and composite objects.
type Graphic interface {
	Draw() string
}

// Composite
// The composite contains a collection of leaf and composite objects.
// It implements the same interface as the leaf struct.
// It delegates the operations to its children.
type Canvas struct {
	Shapes []Graphic
}

// Composite implementation
// The composite struct implements the same interface as the leaf struct.
// This will draw all the shapes (leafs and composites) in the canvas.
func (c *Canvas) Draw() string {
	res := ""
	for _, shape := range c.Shapes {
		res += shape.Draw()
	}
	return res
}

// Leafs
// The leaf struct implements the same interface as the composite struct.
// It represents the end objects in the tree structure.
type (
	Circle struct{}
	Square struct{}
)

// Leaf implementation
// The leaf struct implements the same interface as the composite struct.
func (c Circle) Draw() string {
	return "()"
}
func (s Square) Draw() string {
	return "[]"
}

// Test Composite
// The test function creates a composite object and adds leaf objects to it.
func TestComposite() {
	leaf1 := Circle{}                                    // Leaf
	leaf2 := Square{}                                    // Leaf
	composite := Canvas{Shapes: []Graphic{leaf1, leaf2}} // Composite
	fmt.Println(composite.Draw())                        // Output: ()[]
}
