// Styleguide
// This file will explain the convention of Go language, the rules of writing Go code, and the best practices.
// The Go programming language is designed to be simple and efficient, and its conventions reflect that philosophy.
// The conventions are not just about syntax, but also about how to structure your code, how to name your variables
// and functions, and how to organize your packages.
// Go Style Guide: https://google.github.io/styleguide/go

package styleguide

import (
	"errors"
)

// Modules
// Use only lowercase letters.
// Use the repo pattern for shared projects: "github.com/<username>/<repo>"
var _ = "github.com/john/myproject"
var _ = "myproject"

// Directories
// Use only lowercase letters.
var _ = "httphandler"

// Files
// Use only lowercase letters.
// Use the Go extension: ".go".
var _ = "httphandler.go"

// Singular Names
// Use singular names for specific files like main.go, repository.go, config.go, etc.
var _ = "main.go"
var _ = "repository.go"

// Plural Names
// Use plural names for data collections, slices, models, enums, etc.
var _ = "models.go"
var _ = "enums.go"

// Test Files
// Use only lowercase letters.
// Use the "_test" suffix: "<name>_test.go"
var _ = "httphandler_test.go"

// Packages
// Use only lowercase letters.
// Use the same name as the directory name.
var _ = "httphandler"

// Acronyms
// Use uppercase letters.
var ID = 1
var HTTPClient = 1

// Indicators
// Use Prefixes to indicate the context of a function/variable.
const RoleAdmin = "admin"

// Keywords
// Use the underscore "_" suffix for reserved names.
var Type_ = 1

// Unused
// Use the underscore "_" for unused resources.
var _ = 1

// Variables
// Use camelCase/PascalCase.
var PersonID = 1

// Constants
// Use camelCase/PascalCase.
const EmployeeID = 1

// Functions
// Use camelCase/PascalCase.
func PerformOperation() {}

// New Type
// Use camelCase/PascalCase.
type Number int

// Type Alias
// Use camelCase/PascalCase.
type Decimal = float64

// Interfaces
// Use camelCase/PascalCase.
type Reader interface{ Read() }

// Errors
// Use camelCase/PascalCase.
// Use the "Err" prefix.
// Use lowercase for the error message.
var ErrNotFound = errors.New("not found")

// Structs
// Use camelCase/PascalCase.
type Person struct{ name string }

// Constructors
// Use camelCase/PascalCase.
// Use the "New" prefix.
// Use a pointer as the return type.
func NewPerson(name string) *Person { return &Person{name: name} }

// Methods
// Use camelCase/PascalCase.
// Use a pointer to the receiver type.
// Use a single letter for the receiver name.
func (p *Person) JobName() {}

// Getters
// Use camelCase/PascalCase.
// Don't use the "Get" prefix.
func (p *Person) Name() string { return p.name }

// Setters
// Use camelCase/PascalCase.
// Use the "Set" prefix.
func (p *Person) SetName(name string) { p.name = name }

// Special Names
// Use the following special variable names for specific purposes:
var _ = `
err     // Error values
i, j, k // Loop indexes (especially in nested loops)
n       // Count or length
k       // Map key
v       // Map value
ch      // Channel
res     // Result of a function or operation
buf     // Buffer (e.g., for bytes.Buffer or temporary data)
tmp     // Temporary value or variable
ctx     // Context (context.Context)
tmp     // Temporary value
r, w    // Reader, Writer (w: Response, r: Request)
m       // Generic map
x, y, z // Math / Coordinates
fn      // Function variable
_       // Blank identifier (to ignore a value)
`

// Documentation
// Use single line comments for documentation.
var _ = `
// Takes two integers and returns their sum.
func Sum(x, y int) int {
    ...
}
`

// Section Comment
// Use multi-line comments to create section comments.
var _ = `
/*
   Section Title
*/
`
