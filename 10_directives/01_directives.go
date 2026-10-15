// Directives
// Directives are special comments that provide instructions to the Go compiler or the Go toolchain.
// They are not part of the Go language itself, but they can be used to control the behavior of the compiler
// and the Go toolchain.

package directives

// Go Directives
// Go directives are special comments that provide instructions to the Go compiler or the Go toolchain.
// They follow the format "//go:<name>".
// Go provides a set of directives that can be used in a low-level way for specialized purposes.
// The most common directives are "embed" and "generate".
var _ = `
go:embed foo.txt
go:generate go run foo.go
`

// Other Go Directives
// Go has other directives that are used for specific purposes, such as "cgo" and "testdata".
// These directives are used to control the behavior of the Go compiler and the Go toolchain in specific situations.
var _ = `
go:noescape
go:uintptrescapes
go:noinline
go:norace
go:nosplit
go:linkname localname [importpath.name]
go:wasmimport importmodule importname
`

// Line Directives
// Line directives typically appear in machine-generated code, so that compilers and debuggers will report
// positions in the original input to the generator.
// Below there is an example of a line directive.
var _ = `
line foo.go:10
`
