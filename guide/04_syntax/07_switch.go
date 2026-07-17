// Switch
// The "switch" statement is used to execute different blocks of code based on the value of an expression.
// It provides a cleaner and more readable alternative to multiple "if-else" conditions.
// Syntax:
// - switch {...}
// - switch <variable> {...}
// - switch <declarations>; <variable> {{...}

package syntax

import "fmt"

// Switches
// Switches are used to execute different blocks of code based on the value of an expression.
// They provide a cleaner and more readable alternative to multiple "if-else" conditions.
func Switches() {
	// Basic Switch
	// A standard switch statement evaluates "x" and executes the matching case.
	// Unlike other languages, Go does not require a "break" statement to prevent fallthrough.
	// Output: x == 2
	x := 2
	switch x {
	case 1:
		fmt.Println("x == 1")
	case 2:
		fmt.Println("x == 2")
	}

	// Switch with Default Case
	// The "default" case executes when no other case matches.
	// Output: x == ?
	x = 2
	switch x {
	case 1:
		fmt.Println("x == 1")
	default:
		fmt.Println("x == ?")
	}

	// Case with Multiple Values/Expressions
	// A single case can have multiple values/expressions, acting as an OR condition.
	// It executes the block if any of the expressions match.
	// Output: x == 1 || x == 2
	x = 2
	switch x {
	case 1, 2:
		fmt.Println("x == 1 || x == 2")
	default:
		fmt.Println("?")
	}

	// Switch with Expressions
	// Cases can evaluate expressions dynamically, similar to "if" conditions.
	// Output: x == y + 1
	x = 2
	y := 1
	switch x {
	case y - 1:
		fmt.Println("x == y - 1")
	case y + 1:
		fmt.Println("x == y + 1")
	}

	// Type Switch
	// A type switch is used to compare the type of an interface value.
	// It is similar to a type assertion, but with a switch statement.
	// The "any" type is equivalent to interface{}.
	// Output: z is a float64
	var z any = 3.14
	switch z.(type) {
	case int:
		fmt.Println("z is an int")
	case float64:
		fmt.Println("z is a float64")
	}

	// Fallthrough
	// The "fallthrough" keyword forces execution of the next case, even if it does not match.
	// Unlike other languages, Go does not automatically fall through cases.
	// It can be used to execute multiple cases without duplicating code.
	// In the example below, both cases (case 2 and case 3) will be performed.
	// Output: case 2 | case 3
	x = 2
	switch x {
	case 1:
		fmt.Println("case 1")
		fallthrough
	case 2:
		fmt.Println("case 2")
		fallthrough
	case 3:
		fmt.Println("case 3")
	}

	// Switch Without an Expression
	// When no expression is provided, the switch acts like a series of "if-else" syntax.
	// It is a cleaner alternative to multiple "else if" conditions.
	// Output: x == 2
	x = 2
	switch {
	case x == 1:
		fmt.Println("x == 1")
	case x == 2:
		fmt.Println("x == 2")
	default:
		fmt.Println("?")
	}

	// Switch with Initialization
	// A switch statement can include a short variable declaration to initialize a variable.
	// The variable is scoped to the switch block.
	// Output: x == y
	switch y := 2; x {
	case y:
		fmt.Println("x == y")
	default:
		fmt.Println("x != y")
	}
}
