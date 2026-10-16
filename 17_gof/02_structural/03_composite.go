// Composite
// - Composite is a structural design pattern that lets you compose objects into tree structures to
//   represent part-whole hierarchies.
// - Composite lets clients treat individual objects and compositions of objects uniformly.
// - In other words, it allows you to create a tree structure where each node can be either a leaf
//   or a composite.

package structural

import "fmt"

// Interface
// - The interface defines the common operations for both leaf and composite objects.

type Element interface {
	Render() string
}

// Composite
// - The composite contains a collection of leaf and composite objects.
// - It implements the same interface as the leaf struct.
// - It delegates the operations to its children.

type Group struct {
	Elements []Element
}

// Composite implementation
// - The composite struct implements the same interface as the leaf struct.
// - This will draw all the shapes (leafs and composites) in the canvas.

func (c *Group) Render() string {
	res := "<group>\n"
	for _, element := range c.Elements {
		res += element.Render() + "\n"
	}
	res += "</group>"
	return res
}

// Leafs
// - The leaf struct implements the same interface as the composite struct.
// - It represents the end objects in the tree structure.

type (
	Circle struct{}
	Square struct{}
)

// Leaf implementation
// - The leaf struct implements the same interface as the composite struct.

func (c Circle) Render() string {
	return "<circle />"
}
func (s Square) Render() string {
	return "<square />"
}

// Test Composite
// - The test function creates a composite object and adds leaf objects to it.
// - Output:
//   <group>
//       <group>
//           <circle />
//           <square />
//       </group>
//       <group>
//           <circle />
//           <square />
//       </group>
//   </group>

func TestComposite() {
	composite := &Group{Elements: []Element{
		&Group{Elements: []Element{&Circle{}, &Square{}}},
		&Group{Elements: []Element{&Circle{}, &Square{}}},
	}}
	fmt.Println(composite.Render())
}
