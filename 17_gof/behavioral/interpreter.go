// Interpreter
// The Interpreter pattern is a behavioral design pattern that defines a representation for a language's grammar
// along with an interpreter to evaluate sentences in that language.
// In Go, this is achieved by defining a core expression interface and implementing it across terminal and non-terminal structs.

package behavioral

import (
	"fmt"
)

// Protocol
// The interface below defines the core behavior for evaluating and stringifying expressions.
type Expression interface {
	fmt.Stringer
	Evaluate(context *Context) int
}

// Context
// Contains global information (variable states) needed during interpretation.
type Context struct {
	variables map[string]int
}

func NewContext() *Context {
	return &Context{variables: make(map[string]int)}
}

func (c *Context) SetVariable(name string, value int) {
	c.variables[name] = value
}

func (c *Context) GetVariable(name string) int {
	return c.variables[name]
}

// Terminal Expressions
// Structs representing terminal symbols in the grammar (numbers and variables).
type (
	NumberExpression   struct{ number int }
	VariableExpression struct{ name string }
)

// Terminal Implementations
func (n *NumberExpression) Evaluate(context *Context) int {
	return n.number
}
func (n *NumberExpression) String() string {
	return fmt.Sprintf("%d", n.number)
}

func (v *VariableExpression) Evaluate(context *Context) int {
	return context.GetVariable(v.name)
}
func (v *VariableExpression) String() string {
	return v.name
}

// Non-Terminal Expressions
// Structs implementing composite grammar rules (addition and subtraction).
type (
	AddExpression struct{ x, y Expression }
	SubExpression struct{ x, y Expression }
)

// Non-Terminal Implementations
func (a *AddExpression) Evaluate(context *Context) int {
	return a.x.Evaluate(context) + a.y.Evaluate(context)
}
func (a *AddExpression) String() string {
	return fmt.Sprintf("(%s + %s)", a.x, a.y)
}

func (s *SubExpression) Evaluate(context *Context) int {
	return s.x.Evaluate(context) - s.y.Evaluate(context)
}
func (s *SubExpression) String() string {
	return fmt.Sprintf("(%s - %s)", s.x, s.y)
}

// Client Fluent Builder
// Simplifies building the Abstract Syntax Tree (AST) using a fluent API.
type ExpressionEvaluator struct {
	expression Expression
}

func NewExpressionEvaluator(number int) *ExpressionEvaluator {
	return &ExpressionEvaluator{expression: &NumberExpression{number: number}}
}

func (e *ExpressionEvaluator) Plus(expr Expression) *ExpressionEvaluator {
	e.expression = &AddExpression{x: e.expression, y: expr}
	return e
}

func (e *ExpressionEvaluator) PlusNum(number int) *ExpressionEvaluator {
	e.expression = &AddExpression{x: e.expression, y: &NumberExpression{number: number}}
	return e
}

func (e *ExpressionEvaluator) PlusVar(name string) *ExpressionEvaluator {
	e.expression = &AddExpression{x: e.expression, y: &VariableExpression{name: name}}
	return e
}

func (e *ExpressionEvaluator) Minus(expr Expression) *ExpressionEvaluator {
	e.expression = &SubExpression{x: e.expression, y: expr}
	return e
}

func (e *ExpressionEvaluator) MinusNum(number int) *ExpressionEvaluator {
	e.expression = &SubExpression{x: e.expression, y: &NumberExpression{number: number}}
	return e
}

func (e *ExpressionEvaluator) MinusVar(name string) *ExpressionEvaluator {
	e.expression = &SubExpression{x: e.expression, y: &VariableExpression{name: name}}
	return e
}

func (e *ExpressionEvaluator) Expression() Expression {
	return e.expression
}

func (e *ExpressionEvaluator) Evaluate(context *Context) int {
	return e.expression.Evaluate(context)
}

func (e *ExpressionEvaluator) String() string {
	return e.expression.String()
}

// Test Interpreter
// The test function demonstrates how to build and interpret complex grammar expressions.
func TestInterpreter() {
	expression := NewExpressionEvaluator(5).
		PlusNum(3).
		PlusVar("x").
		Plus(NewExpressionEvaluator(4).
			PlusNum(7).
			MinusVar("y").
			Plus(NewExpressionEvaluator(2).
				PlusNum(3).
				MinusNum(4).
				PlusVar("x").
				Expression(),
			).
			PlusNum(12).
			Expression(),
		)

	context := NewContext()
	context.SetVariable("x", 3)
	context.SetVariable("y", 6)

	fmt.Println(expression)                   // Output: (((5 + 3) + x) + ((((4 + 7) - y) + (((2 + 3) - 4) + x)) + 12))
	fmt.Println(expression.Evaluate(context)) // Output: 32
}
