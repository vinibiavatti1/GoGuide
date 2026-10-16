// String Conversor (strconv)
// - The "strconv" package implements conversions to and from string representations of basic data
//   types.

package library

import (
	"fmt"
	"strconv"
)

//==================================================================================================
// Atoi & Itoa
//==================================================================================================

func AtoiAndItoa() {

	// Atoi
	// - Atoi converts an ASCII string to an int.
	// - Output: 42 (int)

	n, _ := strconv.Atoi("42")
	fmt.Println(n)

	// Itoa
	// - Itoa converts an int to an ASCII string.
	// - Output: 42 (string)

	s := strconv.Itoa(42)
	fmt.Println(s)
}

//==================================================================================================
// Parse Functions
//==================================================================================================

func ParseFunctions() {

	// Parse Functions
	// - Parses string representations into their respective typed values (bool, int, uint, float,
	//   complex).
	// - Output: true | 42 | 42 | 3.14 | (1+2i)

	b, _ := strconv.ParseBool("true")
	i, _ := strconv.ParseInt("42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	f, _ := strconv.ParseFloat("3.14", 64)
	c, _ := strconv.ParseComplex("1+2i", 128)
	fmt.Println(b, i, u, f, c)
}

//==================================================================================================
// Format Functions
//==================================================================================================

func FormatFunctions() {

	// Format Functions
	// - Converts basic types (bool, int, uint, float, complex) into their string representations.
	// - Output: true | 42 | 42 | 3.14 | (1.00+2.00i)

	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatInt(42, 10)
	s3 := strconv.FormatUint(42, 10)
	s4 := strconv.FormatFloat(3.14, 'f', 2, 64)
	s5 := strconv.FormatComplex(1+2i, 'f', 2, 128)
	fmt.Println(s1, s2, s3, s4, s5)
}

//==================================================================================================
// Append Functions
//==================================================================================================

func AppendFunctions() {

	// Append Functions
	// - Appends string forms of values, quotes, or ASCII/Graphic escapes to an existing byte slice
	//   ([]byte).
	// - Output: true 42 3.14 "hello" "caf\u00e9" "hello" 'a' '\u00e9' 'a'

	buf := []byte("")
	buf = strconv.AppendBool(buf, true)
	buf = strconv.AppendInt(buf, 42, 10)
	buf = strconv.AppendFloat(buf, 3.14, 'f', 2, 64)
	buf = strconv.AppendQuote(buf, "hello")
	buf = strconv.AppendQuoteToASCII(buf, "café")
	buf = strconv.AppendQuoteToGraphic(buf, "hello")
	buf = strconv.AppendQuoteRune(buf, 'a')
	buf = strconv.AppendQuoteRuneToASCII(buf, 'é')
	buf = strconv.AppendQuoteRuneToGraphic(buf, 'a')
	fmt.Println(string(buf))
}

//==================================================================================================
// Quote Functions
//==================================================================================================

func QuoteFunctions() {

	// Quote Functions
	// - Returns quoted string or rune literals safely escaped with Go syntax.
	// - Output: "hello" | "hello" | "caf\u00e9" | "hello" | 'a' | '\u00e9' | 'a'

	q1 := strconv.Quote("hello")
	q2, _ := strconv.QuotedPrefix("\"hello\" world")
	q3 := strconv.QuoteToASCII("café")
	q4 := strconv.QuoteToGraphic("hello")
	q5 := strconv.QuoteRune('a')
	q6 := strconv.QuoteRuneToASCII('é')
	q7 := strconv.QuoteRuneToGraphic('a')
	fmt.Println(q1, q2, q3, q4, q5, q6, q7)
}
