// Stringer
// Stringer is a tool to automate the creation of methods that satisfy the fmt.Stringer interface.
// Given the name of a (signed or unsigned) integer type T that has constants defined, stringer will create a new
// self-contained Go source file implementing "func (t T) String() string".
// The file is created in the same package and directory as the package that defines T.
// It has helpful defaults designed for use with go generate.

package tools

// Setup
// The stringer tool is part of the Go tools package and can be installed using the command below.
// It will be installed in the $GOPATH/bin directory, which should be included in your PATH environment variable.
var _ = `
go install golang.org/x/tools/cmd/stringer@latest
`

// Type
// We will define a type named Color, which is an integer type.
// The stringer tool will be used to generate a string representation of the Color type.
type Color int

// Enum
// The Color enumeration defines three colors: Red, Green, and Blue.
// All colors will be considered by the stringer tool to generate a string representation of the enumeration values.
const (
	Red Color = iota
	Green
	Blue
)

// Stringer
// Now, to generate the string representation of the Color enumeration values, we will use the stringer tool.
// The generated code will be placed in a file named color_string.go in the same package as the original file.
// The command to generate the code is as follows:
var _ = `
go generate stringer -type=Color
`

// Directive
// We can also use the //go:generate directive to automate the code generation process.
//
//go:generate stringer -type=Color
