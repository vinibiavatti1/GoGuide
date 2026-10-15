// String
// - A string is a sequence of characters (runes) enclosed in double quotes.
// - Strings are immutable, meaning their values cannot be changed after they are created.
// - We can manipulate strings to create new strings.

package datatypes

import "fmt"

//==================================================================================================
// Declaring Strings
//==================================================================================================

func DeclaringStrings() {

	// Declaring a String
	// - Output: x: Hello World!

	x := "Hello World!"
	fmt.Println("x:", x)

	// Declaring a Multi-line String
	// - We can use the backticks (`) to create a multi-line string.
	// - Output: x: Lorem ipsum dolor sit amet...

	x = `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed et aliquam 
		magna, eu dignissim ante. Donec non semper lectus. Vivamus vel efficitur 
		odio. Integer eu pulvinar augue.
	`
	fmt.Println("x:", x)

	// Declaring a Raw String Literal
	// - We can also use backticks (`) to prevent escape characters (like \n or \t) from being
	//   interpreted.
	// - This is highly used for Regular Expressions and file paths.
	// - Output: x: \d+\s+\w+\t\n

	x = `\d+\s+\w+\t\n`
	fmt.Println("x:", x)

	// Declaring a String using Hexadecimal Notation
	// - We can use the "\x<n>" syntax to define hexadecimal characters.
	// - Output: x: A

	x = "\x41"
	fmt.Println("x:", x)

	// Declaring a String using Octal Notation
	// - We can use the "\<n>" syntax to define octal characters.
	// - Output: x: A

	x = "\101"
	fmt.Println("x:", x)

	// Concatenating Strings
	// - We can create other strings by joining two or more strings.
	// - The "+" operator is used to concatenate strings.
	// - Output: x: Hi, John

	x = "Hi, " + "John"
	fmt.Println("x:", x)

	// Escaping Strings
	// - Some runes must be escaped to be used in a string.
	// - Output: Escape "Example"

	x = "\a" // Alert or bell
	x = "\"" // Escape Double quote (String)
	x = "\\" // Backslash
	x = "\n" // New line
	x = "\t" // Tab
	x = "\r" // Carriage return
	x = "\b" // Backspace
	x = "\f" // Form feed
	x = "\v" // Vertical tab
	fmt.Println("Escape \"Example\"")
}

//==================================================================================================
// String Operations
//==================================================================================================

func StringOperations() {
	x := "Hello, World!"

	// Retrieving String Length
	// - We can use the "len()" function to get the length of a string.
	// - Output: Length: 13

	l := len(x)
	fmt.Println("Length:", l)

	// String Indexing
	// - We can access a character in a string using indexing.
	// - Note: Since the value is a byte, we need to convert it to a string to get the ASCII
	//   representation.
	// - Syntax: x[<index>]
	// - Output: 72 | H

	fmt.Println(x[0], string(x[0]))

	// String Slicing
	// - We can access a substring in a string using slicing.
	// - Note: The end index is exclusive.
	// - Syntax: x[<start>:<end>]

	fmt.Println(x[:5])  // Output: Hello
	fmt.Println(x[7:])  // Output: World!
	fmt.Println(x[2:4]) // Output: ll
	fmt.Println(x[:])   // Output: Hello, World! (Same reference)
	fmt.Println(x[0:0]) // Output: (empty string)

	// Iterating Strings (Using standard for)
	// - We can use the standard for loop to iterate over a string
	// - Output: char: H e l l o...

	for i := 0; i < len(x); i++ {
		c := string(x[i])
		fmt.Println("char:", c)
	}

	// Iterating Strings (Using for-range)
	// - The for-range loop is a more concise way to iterate over a string
	// - Output: char: H e l l o...

	for _, c := range x {
		fmt.Println("char:", string(c))
	}
}
