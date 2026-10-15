// Visitor
// The Visitor pattern is a design pattern that lets you separate algorithms from the objects on which they operate.
// It allows you to add new operations to existing object structures without modifying the structures themselves.

package behavioral

import "fmt"

// Element
// The Element interface declares an "Accept" method that takes a visitor as an argument.
// The visitor will implement the logic for processing the element.
type Shape interface {
	Accept(v Visitor) string
}

// Visitor Interface
// The Visitor interface declares a set of visiting methods for each concrete element class.
// The visitor will implement the logic for processing each element type.
type Visitor interface {
	VisitDot(d *Dot) string
	VisitCircle(c *Circle) string
	VisitRect(r *Rect) string
}

// Visitor
// The Concrete Visitor implements the Visitor interface and provides the logic for processing each element type.
// In this case, we have a visitor that exports the shapes to SVG format.
type SVGExportVisitor struct{}

// Visitor Implementation
// The visitor methods are implemented to handle each shape type.
func (s *SVGExportVisitor) VisitDot(d *Dot) string {
	return "<dot/>"
}
func (s *SVGExportVisitor) VisitCircle(c *Circle) string {
	return "<circle/>"
}
func (s *SVGExportVisitor) VisitRect(r *Rect) string {
	return "<rect/>"
}

// Concrete Elements
// The concrete elements implement the "Accept" method, which calls the appropriate visitor method based on the
// element type.
// Each shape type is represented by a struct that implements the Shape interface.
// The Accept method is implemented to call the corresponding visitor method.
type (
	Dot    struct{}
	Circle struct{}
	Rect   struct{}
)

// Concrete Element Implementation
// The concrete elements implement the "Accept" method, which calls the appropriate visitor method based on the
// element type.
func (d *Dot) Accept(v Visitor) string {
	return v.VisitDot(d)
}
func (c *Circle) Accept(v Visitor) string {
	return v.VisitCircle(c)
}
func (r *Rect) Accept(v Visitor) string {
	return v.VisitRect(r)
}

// Test Visitor
// The function demonstrates how the visitor pattern can be used to process different shapes with a single visitor.
func TestVisitor() {
	d := &Dot{}
	c := &Circle{}
	r := &Rect{}
	exp := &SVGExportVisitor{}
	fmt.Println("Dot:", d.Accept(exp))    // Output: <dot/>
	fmt.Println("Circle:", c.Accept(exp)) // Output: <circle/>
	fmt.Println("Rect:", r.Accept(exp))   // Output: <rect/>
}
