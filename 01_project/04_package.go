// Package
// A package is a collection of related Go files that are compiled together.
// The main concept of a package is to group related code together.
// In Go we only import packages (group of files), not files.
// To import a package, we use the "import" statement.

package project

// Declaring a Package
// To declare a package, we use the "package" statement.
// The package statement must be the first line in the file.
// Usually, the package name is the same as the directory name.
var _ = `
package user
`

// Main Package
// The main package is the entry point of a Go program.
// When the application is run, the main package is searched for the main function.
var _ = `
package main
`

// Importing a Single Package
// To import a single package, we use the "import" statement.
// The import statement must be after the package statement.
// The imported package will be referenced by its name:
// Ex: fmt.Println().
var _ = `
import "fmt"
`

// Importing Multiple Packages
// To import multiple packages, we use parentheses.
var _ = `
import (
	"fmt"
	"math"
)
`

// Importing Packages with Aliases
// We can define an alias for a package.
// This alias will replace the package name, and the alias will be used to reference the package.
var _ = `
import (
	m "math"
)
`

// Importing Packages with Dot
// We can import a package with a dot.
// This will allow us to use the package's exported identifiers without a prefix:
// Ex: Println() instead of fmt.Println().
var _ = `
import (
	. "fmt"
)
`

// Importing Packages with Blank Identifier
// We can import a package with a blank identifier.
// This will allow us to execute the package's init() function without using the package.
var _ = `
import (
	_ "math"
)
`

// Importing Application Packages
// To import the application packages, we have to use the mod.go module name as base.
// Let's suppose we have a module named "github.com/user/app/directory/math".
// To import the math package, we have to use the following import statement:
// Note: Even though the directory has no package definition, the package name is the same as the directory name.
var _ = `
import "github.com/user/app/directory/math"
`
