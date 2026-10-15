// Comments
// Go supports single-line and multi-line comments.
// Comments can be used as GoDoc (documentation) to describe the purpose of:
// - Packages  (First line of a package)
// - Functions (Before the function declaration)
// - Variables (Before the variable declaration)
// - Constants (Before the constant declaration)
// - Types     (Before the type declaration)

package syntax

// Single Line Comment
// Single line comments are used to describe the purpose of a line of code.
// They are also used to document the purpose of a package, function, variable, constant, or type.
var _ = `
// Lorem ipsum dolor sit amet...
`

// Multi Line Comment
// Other comment alternative to document Go resources.
// It can also be used to create "section comments"
var _ = `
/*
   Lorem ipsum dolor sit amet...
*/
`

// Section Comment
// We can use a multi-line comment as a header for the section.
var _ = `
/*
   Section Title
*/
`

// Documentation (GoDoc)
// Go resources (functions, types, consts, vars, packages, ...) should have a documentation when needed.
// Documentation is useful to give details and examples about some Go resource to developers.
// As a convention, documentation should be created with single-line comments.
var _ = `
// Takes two integers and returns their sum.
func Sum(x, y int) int {
    ...
}
`

// Package Documentation
// In Go, a package can span multiple files. To avoid redundancy, the package-level
// documentation should be included only once, typically in the primary file of the package.
// For other files in the package, if file-specific documentation is needed, it should be
// placed at the top of the file, followed by a blank line before the `package` declaration,
// according to Go conventions.
var _ = `
File: math.go
// Documentation...
package math
`
var _ = `
File: sum.go
// Documentation...

package math
`

// New Line
// To create new lines in documentation, we can use an empty comment line.
// This is useful to separate sections in the documentation.
var _ = `
// Takes two integers and returns their sum.
//
// See math package for more details.
func Sum(x, y int) int {
    ...
}
`

// Deprecated
// Deprecated comments are used to mark a function, variable, constant, or type as deprecated.
// To mark a function, variable, constant, or type as deprecated, we can use the `// Deprecated:` comment.
// Note: The annotation must be placed in a new line.
var _ = `
// Takes two integers and returns their sum.
//
// Deprecated: Use Add instead.
func Sum(x, y int) int {
    ...
}
`
