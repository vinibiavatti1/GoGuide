// FMT (Formatting & I/O)
// - The "fmt" package implements formatted I/O with functions analogous to C's printf and scanf.
// - Print: Used to format and write output to standard output, strings, io.Writers, or byte slices.
// - Scan: Used to read and parse input from standard input, strings, or io.Readers.

package basics

import (
	"bytes"
	"fmt"
)

//==================================================================================================
// Print Functions
//==================================================================================================

func PrintFunctions() {

	// Print
	// - Prints the operands to standard output using their default formats.
	// - Output: Hello World

	fmt.Print("Hello World")

	// Println
	// - Prints the operands to standard output with spaces between them, appending a newline at the
	//   end.
	// - Output: Hello World\n

	fmt.Println("Hello World")

	// Printf
	// - Prints formatted text to standard output according to a format specifier.
	// - Output: Value: 42.00

	fmt.Printf("Value: %.2f", 42.0)
}

//==================================================================================================
// SPrint Functions
//==================================================================================================

func SPrintFunctions() {

	// Sprint
	// - Formats using the default formats and returns the resulting string.
	// - Output: Hello World

	x := fmt.Sprint("Hello World")
	fmt.Println(x)

	// Sprintln
	// - Formats using the default formats, appending spaces between operands and a newline at the
	//   end, returning a string.
	// - Output: Hello World\n

	y := fmt.Sprintln("Hello World")
	fmt.Println(y)

	// Sprintf
	// - Formats according to a format specifier and returns the resulting string.
	// - Output: Value: 42.00

	z := fmt.Sprintf("Value: %.2f", 42.0)
	fmt.Println(z)
}

//==================================================================================================
// Fprint Functions
//==================================================================================================

func FPrintFunctions() {

	// Fprint
	// - Formats using the default formats and writes to standard output or any io.Writer.
	// - Output: Hello World

	var b1 bytes.Buffer
	fmt.Fprint(&b1, "Hello World")
	fmt.Println(b1.String())

	// Fprintln
	// - Formats using the default formats, appending spaces and a newline, writing to an io.Writer.
	// - Output: Hello World\n

	var b2 bytes.Buffer
	fmt.Fprintln(&b2, "Hello World")
	fmt.Println(b2.String())

	// Fprintf
	// - Formats according to a format specifier and writes the result to an io.Writer.
	// - Output: Value: 42.00

	var b3 bytes.Buffer
	fmt.Fprintf(&b3, "Value: %.2f", 42.0)
	fmt.Println(b3.String())
}

//==================================================================================================
// Scan
//==================================================================================================

func ScanFunctions() {

	// Scan
	// - The "fmt.Scan" function reads input from standard input, separated by spaces.
	// - Output: <input>

	var x string
	fmt.Print("Enter your name: ")
	fmt.Scan(&x)
	fmt.Println(x)

	// Scanln
	// - The "fmt.Scanln" function reads input until a newline character is encountered.
	// - Output: <input>

	var y string
	fmt.Print("Enter your full name: ")
	fmt.Scanln(&y)
	fmt.Println(y)

	// Scanf
	// - The "fmt.Scanf" function reads input from standard input according to a format specifier.
	// - Output: <input>

	var z int
	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &z)
	fmt.Println(z)
}

//==================================================================================================
// Sscan
//==================================================================================================

func SscanFunctions() {

	// Sscan
	// - The "fmt.Sscan" function reads space-separated values from a string into variables.
	// - Output: John 30

	var x1 string
	var y1 int
	fmt.Sscan("John 30", &x1, &y1)
	fmt.Println(x1, y1)

	// Sscanln
	// - The "fmt.Sscanln" function reads space-separated values from a string, stopping at a
	//   newline.
	// - Output: John 30

	var x2 string
	var y2 int
	fmt.Sscanln("John\n30", &x2, &y2)
	fmt.Println(x2, y2)

	// Sscanf
	// - The "fmt.Sscanf" function reads values from a string according to a format specifier.
	// - Output: John 30

	var x3 string
	var y3 int
	fmt.Sscanf("John 30", "%s %d", &x3, &y3)
	fmt.Println(x3, y3)
}

//==================================================================================================
// Fscan
//==================================================================================================

func FscanFunctions() {

	// Fscan
	// - The "fmt.Fscan" function reads space-separated values from an io.Reader (like a buffer or
	//   file) into variables.
	// - Output: John 30

	b1 := bytes.NewBufferString("John 30")
	var x1 string
	var y1 int
	fmt.Fscan(b1, &x1, &y1)
	fmt.Println(x1, y1)

	// Fscanln
	// - The "fmt.Fscanln" function reads space-separated values from an io.Reader, stopping at a
	//   newline.
	// - Output: John 30

	b2 := bytes.NewBufferString("John\n30")
	var x2 string
	var y2 int
	fmt.Fscanln(b2, &x2, &y2)
	fmt.Println(x2, y2)

	// Fscanf
	// - The "fmt.Fscanf" function reads values from an io.Reader according to a format specifier.
	// - Output: John 30

	b3 := bytes.NewBufferString("John 30")
	var x3 string
	var y3 int
	fmt.Fscanf(b3, "%s %d", &x3, &y3)
	fmt.Println(x3, y3)
}

//==================================================================================================
// Append Functions
//==================================================================================================

func AppendFunctions() {

	// Append
	// - Appends the result of formatting to a byte slice and returns the updated slice.
	// - Output: Say: Hello

	buf1 := []byte("Say: ")
	buf1 = fmt.Append(buf1, "Hello")
	fmt.Println(string(buf1))

	// Appendln
	// - Appends the result of formatting with spaces and a newline to a byte slice.
	// - Output: Msg: World\n

	buf2 := []byte("Msg: ")
	buf2 = fmt.Appendln(buf2, "World")
	fmt.Println(string(buf2))

	// Appendf
	// - Appends the result of formatting according to a format specifier to a byte slice.
	// - Output: Value: 42.00

	buf3 := []byte("Value: ")
	buf3 = fmt.Appendf(buf3, "%.2f", 42.0)
	fmt.Println(string(buf3))
}

//==================================================================================================
// Formats
//==================================================================================================

func Formats() {

	// Formats
	// - Demonstrates the formatting verbs available in the "fmt" package.
	// - Formats:
	//   %v   - the value in a default format
	//   %+v  - val plus field names (for structs)
	//   %#v  - a Go-syntax representation of the value
	//   %T   - a Go-syntax representation of the type of the value
	//   %%   - a literal percent sign
	//   %t   - word true or false (boolean)
	//   %b   - base 2 (integer)
	//   %c   - the character represented by the corresponding Unicode code point
	//   %d   - base 10 (integer)
	//   %o   - base 8 (integer)
	//   %O   - base 8 with 0o prefix
	//   %x   - base 16, with lower-case letters for a-f
	//   %X   - base 16, with upper-case letters for A-F
	//   %U   - Unicode format: U+1234; same as "U+%04X"
	//   %f   - decimal point without exponent (e.g., 123.456)
	//   %e   - scientific notation (e.g., -1.234456e+78)
	//   %s   - uninterpreted bytes of a string or slice
	//   %q   - a double-quoted string safely escaped with Go syntax
	//   %p   - base 16 notation, with leading 0x (pointer)

	type sample struct{ Name string }
	s := sample{Name: "Go"}
	n := 42

	fmt.Printf("%v", s)     // Output: {Go}
	fmt.Printf("%+v", s)    // Output: {Name:Go}
	fmt.Printf("%#v", s)    // Output: basics.sample{Name:"Go"}
	fmt.Printf("%T", s)     // Output: basics.sample
	fmt.Printf("%%")        // Output: %
	fmt.Printf("%t", true)  // Output: true
	fmt.Printf("%b", n)     // Output: 101010
	fmt.Printf("%c", 65)    // Output: A
	fmt.Printf("%d", n)     // Output: 42
	fmt.Printf("%o", n)     // Output: 52
	fmt.Printf("%O", n)     // Output: 0o52
	fmt.Printf("%x", 255)   // Output: ff
	fmt.Printf("%X", 255)   // Output: FF
	fmt.Printf("%U", 'A')   // Output: U+0041
	fmt.Printf("%f", 3.14)  // Output: 3.140000
	fmt.Printf("%e", 123.4) // Output: 1.234000e+02
	fmt.Printf("%s", "hi")  // Output: hi
	fmt.Printf("%q", "hi")  // Output: "hi"
	fmt.Printf("%p", &n)    // Output: 0xc0000a6018
}
