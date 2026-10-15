// Functions
// - Functions in Go are defined using the "func" keyword.
// - They can take zero or more parameters and return zero or more values.
// - Syntax: func <name>(<parameters> <type>...) <return> {...}

package functions

import "fmt"

// Private Functions
// - Functions with the name starting with a lowercase letter are private functions.
// - They can only be called inside the package where they are defined.
func privateFunction() {
}

// Public Functions
// - Functions with the name starting with an uppercase letter are public functions.
// - They can be called from other packages.
// - The function below can call the private function, since it is located at the same package.
func PublicFunction() {
}

// Parameters
// - Functions can take zero or more parameters.
// - Parameters are defined with a name and a type.
func FunctionParameters(x int, y string) {
}

// Grouped Parameters
// - Multiple parameters of the same type can be grouped together.
func GroupedParameters(x, y, z int) {
}

// Return Value
// - Functions can return zero or more values.
// - Return values are defined after the parameter list.
func FunctionReturn() int {
	return 1
}

// Multiple Return Values
// - Functions can return multiple values.
// - The return values are separated by commas and enclosed in parentheses.
func FunctionMultipleReturn() (int, int) {
	return 1, 2
}

// Named Return Values (Naked Returns)
// - Functions can return named values.
// - Named return values are initialized with the zero value of their type.
// - An empty return statement (naked return) returns the named values implicitly.
func NamedReturnValues() (x, y int) {
	x, y = 1, 1
	return
}

// Variadic Parameters
// - Functions can take a variable number of arguments using variadic parameters.
// - The variadic parameter is a slice that contains all arguments passed to the function.
// - To define a variadic parameter, use the "..." operator followed by the type.
// - Output: [1 2 3]
func VariadicParameters(nums ...int) {
	fmt.Println(nums)
}

// Function As Parameter
// - Go supports higher-order functions (HOFs), which means that functions can be passed as
//   arguments, returned from other functions, and assigned to variables.
// - This is useful for callbacks and deferred execution.
func FunctionAsParameter(fn func()) {
	fn()
}

// Function As Return Type
// - Functions can be returned as values from other functions.
// - This is useful for creating closures and higher-order functions.
func FunctionAsReturnType() func() {
	return func() {}
}

// Unused Function
// - Unused functions can be declared using the blank identifier "_".
func _() {
}
