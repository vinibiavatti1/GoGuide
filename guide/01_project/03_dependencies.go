// Dependencies
// Dependencies are external packages that are used in the project.
// Dependencies are managed by the Go module system, which is a built-in dependency management system in Go.
// The Go module system allows you to define and manage dependencies for your Go projects easily.
// It also allows you to work with multiple versions of a dependency and to share your code with others easily.

package project

// Adding Dependencies
// Dependencies are registered in the go.mod file.
// To add a dependency, you can use the go get command.
// The go get command will add the dependency to the go.mod file and download the dependency.
// If version is not specified, the latest version will be used.
var _ = "go get <module-name>@<version>"

// Updating Dependencies
// To update a dependency, we can use the go get command with the -u flag.
// The -u flag will update the dependency to the latest version.
var _ = "go get -u <module-name>"

// Removing Dependencies
// To remove a dependency, we can use the go mod tidy command.
// The go mod tidy command will remove any dependencies that are not used in the module.
var _ = "go mod tidy"

// Installing Dependencies
// Go install all the dependencies in the go.mod file automatically when the application is built or run.
var _ = "go run ."
