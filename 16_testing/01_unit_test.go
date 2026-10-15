// Unit Tests
// Unit tests are used to test individual components of the codebase in isolation.
// They are typically written using the testing package in Go and are run using the go test command.
// Unit Test Conventions:
// - Unit tests are placed in the same package as the code they are testing;
// - Test files should be named with the _test suffix, e.g., mypackage_test.go;
// - Test functions should start with the word Test and take a pointer to testing.T as an argument.

package testing

import (
	"testing"
)

// Creating a Function
// We will create a simple function that adds two integers together.
// This function will be used in our unit test.
func Sum(x, y int) int {
	return x + y
}

// Writing a Unit Test
// Now, we will write a unit test for the Sum function.
// Unit tests have a parameter of type *testing.T.
func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Sum(2, 3) = %d; expected %d", result, expected)
	}
}

// Running the Test
// To run the test, use the go test command in the terminal.
// This will execute all the tests in the current package.
// The command will look at test files (ending with _test.go) and run the functions that start with Test.
var _ = "go test"

// Result Example
// The output of the test will look something like this:
var _ = `
PASS
ok      guide/testing      1.703s
`
