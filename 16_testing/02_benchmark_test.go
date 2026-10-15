// Benchmark Tests
// Benchmark tests are used to measure the performance of a piece of code.
// They are typically written using the testing package in Go and are run using the go test -bench command.
// Benchmark Conventions:
// - Benchmark tests are placed in the same package as the code they are testing;
// - Benchmark test files should be named with the _test suffix, e.g., mypackage_test.go;
// - Benchmark test functions should start with the word Benchmark and take a pointer to testing.B as an argument.

package testing

import (
	"testing"
)

// Creating a Function
// We will create a simple that will be used in our benchmark test.
// This function will concatenate two strings using the + operator.
func Concat(x, y string) string {
	return x + y
}

// Creating a Benchmark Test
// Now, we will write a benchmark test for the ConcatWithPlus and ConcatWithBuilder functions.
// Benchmark tests have a parameter of type *testing.B.
// The "b.Loop()" method is used to run the benchmark multiple times.
// The testing package will automatically determine the number of iterations based on the performance of the code.
func BenchmarkConcat(b *testing.B) {
	for b.Loop() {
		_ = Concat("Hello", "World")
	}
}

// Running the Benchmark
// To run the benchmark test, use the go test -bench command in the terminal.
// This will execute all the benchmark tests in the current package.
// The command will look at test files (ending with _test.go) and run the functions that start with Benchmark.
var _ = "go test -bench ."

// Result Example
// The output of the benchmark test will look something like this:
var _ = `
goos: windows
goarch: amd64
pkg: guide/testing
cpu: 13th Gen Intel(R) Core(TM) i5-1340P
BenchmarkConcatWithPlus-16      120189393                9.958 ns/op
PASS
ok      guide/testing      3.961s
`
