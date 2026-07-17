// Flag
// The flag package provides a simple way to define and parse command-line flags.
// It is designed to be easy to use and understand, while still being powerful enough
// for most use cases.
// The package provides a Flag type that can be used to define flags, and a
// Parse function that can be used to parse command-line arguments.

package library

import (
	"flag"
	"fmt"
)

// Declaring Flags
// We can use the datatype functions to declare flags of the respective type.
// These flags can be defined directly in the global scope or inside a function, as variables.
// The statements below declare flags for -p1 -p2 -p3 -p4 -p5 -p6 -p7 -p8 command line parameters.
var (
	p1Flag = flag.String("p1", "default1", "p1 description...")
	p2Flag = flag.Int("p2", 0, "p2 description...")
	p3Flag = flag.Int64("p3", 0, "p3 description...")
	p4Flag = flag.Uint("p4", 0, "p4 description...")
	p5Flag = flag.Uint64("p5", 0, "p5 description...")
	p6Flag = flag.Bool("p6", false, "p6 description...")
	p7Flag = flag.Float64("p7", 0.0, "p7 description...")
	p8Flag = flag.Duration("p8", 0, "p8 description...")
)

// Parsing Flags
// After defining the flags, we need to call the flag.Parse() function to parse the command-line arguments.
// This function should be called after all flags have been defined, but before any flags are accessed.
func init() {
	flag.Parse()
}

// Accessing Flags
// To access the values of the flags, we can use the variables defined above.
// The variables are pointer variables, so we need to dereference them to get the actual values.
func ProcessFlags() {

	// Acessing Flags
	// We will print all the flags to the console.
	fmt.Println("p1:", *p1Flag)
	fmt.Println("p2:", *p2Flag)
	fmt.Println("p3:", *p3Flag)
	fmt.Println("p4:", *p4Flag)
	fmt.Println("p5:", *p5Flag)
	fmt.Println("p6:", *p6Flag)
	fmt.Println("p7:", *p7Flag)
	fmt.Println("p8:", *p8Flag)
}

// Setting Flags in Command Line
// To set the flags in the command line, the following syntaxs are allowed:
var _ = `
 -flag
--flag   // double dashes are also permitted
 -flag=x
 -flag x // non-boolean flags only
`

// Help Flag
// We can use the -h or --help flag to get help on the flags.
// This will print the usage information for the flags, including the default values and descriptions.
// The help flag is automatically added to the flag set, so we don't need to define it ourselves.
var _ = `
 -h
--help
`
