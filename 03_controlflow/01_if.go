// If
// The "if" keyword is used to execute code based on conditions.
// It allows us to control the flow of execution by evaluating conditions.
// Syntax:
// - if <condition> {...}
// - if <condition> {...} else {...}
// - if <condition> {...} else if {...}
// - if <declarations>; <condition> {...}

package controlflow

import "fmt"

// If With Comparison Operators
// You can use comparison operators (e.g., ==, !=, >, <) in if conditions.
func IfWithComparisonOperators() {
	// Equals Condition
	// We can use the "==" operator to check if two values are equal.
	// Output: x == y
	x, y := 1, 1
	if x == y {
		fmt.Println("x == y")
	}

	// Not Equals Condition
	// We can use the "!=" operator to check if two values are not equal.
	// Output: x != y
	x, y = 1, 2
	if x != y {
		fmt.Println("x != y")
	}

	// Greater Than Condition
	// We can use the ">" operator to check if a value is greater than another.
	// Output: x > y
	x, y = 2, 1
	if x > y {
		fmt.Println("x > y")
	}

	// Greater Than Or Equal To Condition
	// We can use the ">=" operator to check if a value is greater than or equal to another.
	// Output: x >= y
	x, y = 1, 1
	if x >= y {
		fmt.Println("x >= y")
	}

	// Less Than Condition
	// We can use the "<" operator to check if a value is less than another.
	// Output: x < y
	x, y = 1, 2
	if x < y {
		fmt.Println("x < y")
	}

	// Less Than Or Equal To Condition
	// We can use the "<=" operator to check if a value is less than or equal to another.
	// Output: x <= y
	x, y = 1, 1
	if x <= y {
		fmt.Println("x <= y")
	}
}

// If With Logical Operators (&&, || e !)
// Logical operators can be used in if conditions to combine multiple conditions.
func IfWithLogicalOperators() {
	// AND Operator (&&)
	// The AND operator (&&) is used to check if all conditions are true.
	// Output: x == y && x == z
	x, y, z := 1, 1, 1
	if x == y && x == z {
		fmt.Println("x == y && x == z")
	}

	// OR Operator (||)
	// The OR operator (||) is used to check if at least one condition is true.
	// Output: x == y || x == z
	x, y, z = 1, 2, 1
	if x == y || x == z {
		fmt.Println("x == y || x == z")
	}

	// NOT Operator (!)
	// The NOT operator (!) is used to negate a condition.
	// Output: a == false
	a := false
	if !a {
		fmt.Println("a == false")
	}
}

// If-Else
// The "else" keyword is used to execute code when the condition is false.
// It allows us to execute code when the condition is not met.
// The "else if" keyword is used to evaluate multiple conditions.
func IfElse() {
	// If-Else
	// The "else" keyword is used to execute code when the condition is false.
	// Output: x != y
	x, y, z := 1, 2, 1
	if x == y {
		fmt.Println("x == y")
	} else {
		fmt.Println("x != y")
	}

	// Else-If
	// The "else if" keyword is used to evaluate multiple conditions.
	// Output: x == z
	x, y, z = 1, 2, 1
	if x == y {
		fmt.Println("x == y")
	} else if x == z {
		fmt.Println("x == z")
	}

	// Else-If (With Else)
	// The "else" keyword can be used with "else if" to execute code when all conditions are false.
	// Output: x != y && x != z
	x, y, z = 1, 2, 3
	if x == y {
		fmt.Println("x == y")
	} else if x == z {
		fmt.Println("x == z")
	} else {
		fmt.Println("x != y && x != z")
	}

	// Else-If (With Switch)
	// A switch statement without an expression have the same behavior of a "else If"
	// Output: x == z
	x, y, z = 1, 2, 1
	switch {
	case x == y:
		fmt.Println("x == y")
	case x == z:
		fmt.Println("x == z")
	default:
		fmt.Println("x != y && x != z")
	}
}

// If With Declaration
// We can declare variables inside the if condition.
// It allows us to limit the scope of the variable to the if block.
func IfWithDeclarations() {
	// If with Declaration
	// We can declare a variable inside the if condition (Like in for statements).
	// In this case, "x" is only available inside the if block.
	// Output: x == true
	if x := true; x {
		fmt.Println("x == true")
	}

	// If with Multiple Declarations
	// We can declare multiple variables inside the if condition.
	// In this case, "x" and "y" are only available inside the if block.
	// Output: x < y
	if x, y := 1, 2; x < y {
		fmt.Println("x < y")
	}
}
