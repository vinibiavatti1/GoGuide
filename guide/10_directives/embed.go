// Embed
// The embed directive is a compiler directive that allows programs to include arbitrary files and folders in the
// Go binary at build time.
// Go source files that import "embed" can use the the directive to initialize a variable of type string, []byte,
// or FS with the contents of files read from the package directory or subdirectories at compile time.
// Embed directives accept paths relative to the directory containing the Go source file.
// Syntax:
//   //go:embed <path>

package directives

import (
	"embed"
	"fmt"
)

// Embed File in String
// The variable File1 is a string that contains the contents of the file resources/test.txt
//
//go:embed resources/test.txt
var File1 string

// Embed File in Byte Slice
// The variable File2 is a byte slice that contains the contents of the file resources/test.txt
//
//go:embed resources/test.txt
var File2 []byte

// Embed Folder
// The variable Folder is a FileSystem that contains the contents of the folder resources
//
//go:embed resources
var Folder embed.FS

// Test Embed
// The function below demonstrates how to use the embed directive to include files and folders in the Go binary at build time.
func TestEmbed() {

	// File1
	// String containing the contents of the file resources/test.txt
	fmt.Println(File1) // Output: Hello, World!

	// File2
	// Byte slice containing the contents of the file resources/test.txt
	fmt.Println(File2) // Output: [72 101 108 108 111 44 32 87 111 114 108 100 33]

	// Folder
	// FileSystem containing the contents of the folder resources
	// The file resources/test.txt can be accessed using the Open method of the FileSystem
	f, _ := Folder.ReadFile("resources/test.txt")
	fmt.Println(string(f)) // Output: Hello, World!
}
