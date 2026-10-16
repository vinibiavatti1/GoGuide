// Interpreter
// - The Interpreter pattern is a behavioral design pattern that defines a representation for a
//   language's grammar along with an interpreter to evaluate sentences in that language.
// - In Go, this is achieved by defining a core expression interface and implementing it across
//   terminal and non-terminal structs.

package behavioral

import "fmt"

// Protocol
// - The interface below defines the core behavior for evaluating and stringifying expressions.

type Expression interface {
	Evaluate(ctx map[string]float64) float64
	String() string
}

// Expressions
// - VarExp represents a variable terminal expression.
// - ConstExp represents a constant terminal expression.
// - AddExp represents a non-terminal addition expression.
// - MulExp represents a non-terminal multiplication expression.

type VarExp struct {
	Name string
}
type ConstExp struct {
	Value float64
}
type AddExp struct {
	Exp1 Expression
	Exp2 Expression
}
type MulExp struct {
	Exp1 Expression
	Exp2 Expression
}

// Evaluate Implementations
// - Each expression type implements the Evaluate method to compute its value based on the context.

func (v *VarExp) Evaluate(ctx map[string]float64) float64 {
	return ctx[v.Name]
}
func (v *ConstExp) Evaluate(ctx map[string]float64) float64 {
	return v.Value
}
func (v *AddExp) Evaluate(ctx map[string]float64) float64 {
	return v.Exp1.Evaluate(ctx) + v.Exp2.Evaluate(ctx)
}
func (v *MulExp) Evaluate(ctx map[string]float64) float64 {
	return v.Exp1.Evaluate(ctx) * v.Exp2.Evaluate(ctx)
}

// String Implementations
// - Each expression type implements the String method to return its string representation.

func (v *VarExp) String() string {
	return v.Name
}
func (v *ConstExp) String() string {
	return fmt.Sprintf("%.2f", v.Value)
}
func (v *AddExp) String() string {
	return fmt.Sprintf("(%s + %s)", v.Exp1.String(), v.Exp2.String())
}
func (v *MulExp) String() string {
	return fmt.Sprintf("(%s * %s)", v.Exp1.String(), v.Exp2.String())
}

// Test Interpreter
// - The test function demonstrates how to build and interpret complex grammar expressions.
// - Output: (((a * b) + (a * b)) + (2.00 * (a * b)))
// - Output: 100

func TestInterpreter() {
	exp := &AddExp{
		Exp1: &AddExp{
			Exp1: &MulExp{
				Exp1: &VarExp{Name: "a"},
				Exp2: &VarExp{Name: "b"},
			},
			Exp2: &MulExp{
				Exp1: &VarExp{Name: "a"},
				Exp2: &VarExp{Name: "b"},
			},
		},
		Exp2: &MulExp{
			Exp1: &ConstExp{Value: 2},
			Exp2: &MulExp{
				Exp1: &VarExp{Name: "a"},
				Exp2: &VarExp{Name: "b"},
			},
		},
	}
	fmt.Println(exp.String())
	fmt.Println(exp.Evaluate(map[string]float64{"a": 5, "b": 5}))
}
