// Strings
// Strings is a package that provides functions to manipulate strings.
// It includes many utility functions for string operations.
// This document is divided into sections, each demonstrating a different set of functions.
// - Compare Functions
// - Contains Functions
// - Index Functions
// - Replace Functions
// - Cut Functions
// - Split Functions
// - Fields Functions
// - Case Functions
// - Trim Functions
// - Other Functions
// - Types

package library

import (
	"fmt"
	"strings"
	"unicode"
)

// String Compare Functions
// Example of comparison functions in the strings package.
func StringCompareFunctions() {

	// Compare
	// Compare compares two strings lexicographically.
	// It returns:
	// - 0 if a == b
	// - -1 if a < b
	// - +1 if a > b
	cmp := strings.Compare("Hello", "World")
	fmt.Println("cmp:", cmp) // Output: cmp: -1

	// EqualFold
	// EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under Unicode case-folding.
	// It is case-insensitive.
	eq := strings.EqualFold("Hello", "hello")
	fmt.Println("eq:", eq) // Output: eq: true
}

// String Contains Functions
// Example of functions that check for the presence of substrings in strings.
func StringContainsFunctions() {

	// Contains
	// Contains reports whether substr is within s.
	has := strings.Contains("Hello, World!", "World")
	fmt.Println("has:", has) // Output: has: true

	// ContainsAny
	// ContainsAny reports whether any Unicode code points in chars are within s.
	// The example below checks if any vowels are present in the string.
	has = strings.ContainsAny("Hello, World!", "aeiou")
	fmt.Println("has:", has) // Output: has: true

	// ContainsFunc
	// ContainsFunc reports whether any Unicode code points r within s satisfy f(r).
	// The example below checks if any vowels are present in the string using a custom function.
	has = strings.ContainsFunc("Hello, World!", func(r rune) bool {
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	})
	fmt.Println("has:", has) // Output: has: true

	// ContainsRune
	// ContainsRune reports whether the Unicode code point r is within s.
	has = strings.ContainsRune("Hello, World!", 'H')
	fmt.Println("has:", has) // Output: has: true

	// HasPrefix
	// HasPrefix reports whether the string s begins with prefix.
	has = strings.HasPrefix("Hello, World!", "Hello")
	fmt.Println("has:", has) // Output: has: true

	// HasSuffix
	// HasSuffix reports whether the string s ends with suffix.
	has = strings.HasSuffix("Hello, World!", "World!")
	fmt.Println("has:", has) // Output: has: true
}

// String Index Functions
// Example of functions that return the index of substrings in strings.
func StringIndexFunctions() {

	// Index
	// Index returns the index of the first instance of substr in s, or -1 if substr is not present.
	i := strings.Index("Hello, World!", "World")
	fmt.Println("i:", i) // Output: i: 7 (Found 'World' at index 7)

	// IndexAny
	// IndexAny returns the index of the first instance of any Unicode code point in chars within s,
	// or -1 if none are present.
	i = strings.IndexAny("Hello, World!", "aeiou")
	fmt.Println("i:", i) // Output: i: 1 (Found 'e' at index 1)

	// IndexByte
	// IndexByte returns the index of the first instance of c in s, or -1 if c is not present.
	i = strings.IndexByte("Hello, World!", 'W')
	fmt.Println("i:", i) // Output: i: 7 (Found 'W' at index 7)

	// IndexFunc
	// IndexFunc returns the index of the first Unicode code point satisfying f(c) in s,
	// or -1 if none do.
	i = strings.IndexFunc("Hello, World!", func(r rune) bool {
		return r == 'W'
	})
	fmt.Println("i:", i) // Output: i: 7 (Found 'W' at index 7)

	// IndexRune
	// IndexRune returns the index of the first instance of r in s, or -1 if r is not present.
	i = strings.IndexRune("Hello, World!", 'W')
	fmt.Println("i:", i) // Output: i: 7 (Found 'W' at index 7)

	// LastIndex
	// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present.
	i = strings.LastIndex("Hello, World!", "o")
	fmt.Println("i:", i) // Output: i: 8 (Found last 'o' at index 8)

	// LastIndexAny
	// LastIndexAny returns the index of the last instance of any Unicode code point in chars within s,
	// or -1 if none are present.
	i = strings.LastIndexAny("Hello, World!", "aeiou")
	fmt.Println("i:", i) // Output: i: 8 (Found last 'o' at index 8)

	// LastIndexByte
	// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present.
	i = strings.LastIndexByte("Hello, World!", 'o')
	fmt.Println("i:", i) // Output: i: 8 (Found last 'o' at index 8)

	// LastIndexFunc
	// LastIndexFunc returns the index of the last Unicode code point satisfying f(c) in s,
	// or -1 if none do.
	i = strings.LastIndexFunc("Hello, World!", func(r rune) bool {
		return r == 'o'
	})
	fmt.Println("i:", i) // Output: i: 8 (Found last 'o' at index 8)
}

// String Replace Functions
// Example of functions that replace substrings in strings.
func StringReplaceFunctions() {

	// Replace
	// Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
	// If n < 0, all instances are replaced.
	x := strings.Replace("Hello", "l", "1", 1)
	fmt.Println("x:", x) // Output: x: He1lo (Only the first 'l' is replaced)

	// ReplaceAll
	// ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new.
	// It is an alias for Replace(s, old, new, -1).
	x = strings.ReplaceAll("Hello", "l", "1")
	fmt.Println("x:", x) // Output: x: He11o (All 'l' are replaced)
}

// String Cut Functions
// Example of functions that cut strings into two parts.
func StringCutFunctions() {

	// Cut
	// Cut slices s into two parts: the text before the first instance of sep and the text after.
	// It returns the text before sep, the text after sep, and a boolean indicating if sep was found.
	before, after, found := strings.Cut("Hello, World!", ", ")
	fmt.Println("before:", before) // Output: before: Hello
	fmt.Println("after:", after)   // Output: after: World!
	fmt.Println("found:", found)   // Output: found: true

	// CutPreffix
	// CutPrefix returns s without the provided leading prefix string.
	// It returns the string after the prefix and a boolean indicating if the prefix was found.
	after, found = strings.CutPrefix("Hello, World!", "Hello")
	fmt.Println("after:", after) // Output: after: , World!
	fmt.Println("found:", found) // Output: found: true

	// CutSuffix
	// CutSuffix returns s without the provided trailing suffix string.
	// It returns the string before the suffix and a boolean indicating if the suffix was found.
	after, found = strings.CutSuffix("Hello, World!", "World!")
	fmt.Println("after:", after) // Output: after: Hello,
	fmt.Println("found:", found) // Output: found: true
}

// String Split Functions
// Example of functions that split strings into parts.
func StringSplitFunctions() {

	// Split
	// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
	// If sep is empty, Split splits after each UTF-8 sequence.
	// If sep is not found, Split returns a slice of s.
	// If s is empty, Split returns an empty slice.
	parts := strings.Split("a,b,c", ",")
	fmt.Println("parts:", parts) // Output: parts: [a b c]

	// SplitAfter
	// SplitAfter slices s into all substrings after each instance of sep and returns a slice of the substrings.
	parts = strings.SplitAfter("a,b,c", ",")
	fmt.Println("parts:", parts) // Output: parts: [a, b, c] (The separator is included in the result)

	// SplitAfterN
	// SplitAfterN slices s into all substrings after each instance of sep and returns a slice of the substrings.
	// The result is limited to n substrings.
	parts = strings.SplitAfterN("a,b,c", ",", 2)
	fmt.Println("parts:", parts) // Output: parts: [a, b,c] (The separator is included in the result)

	// SplitAfterSeq
	// SplitAfterSeq slices s into all substrings after each instance of sep and returns a slice of the substrings.
	iter := strings.SplitAfterSeq("a,b,c", ",")
	for v := range iter {
		fmt.Println(v) // Output: a, b, c (The separator is included in the result)
	}

	// SplitN
	// SplitN slices s into all substrings separated by sep and returns a slice of the substrings.
	// The result is limited to n substrings.
	parts = strings.SplitN("a,b,c", ",", 2)
	fmt.Println("parts:", parts) // Output: parts: [a b,c] (The separator is not included in the result)

	// SplitSeq
	// SplitSeq slices s into all substrings separated by sep and returns a slice of the substrings.
	iter = strings.SplitSeq("a,b,c", ",")
	for v := range iter {
		fmt.Println(v) // Output: a b c (The separator is not included in the result)
	}
}

// String Fields Functions
// Example of functions that split strings into fields.
func StringFieldsFunctions() {

	// Fields
	// Fields splits s around each instance of one or more consecutive white space characters.
	fields := strings.Fields("  Hello,   World!  ")
	fmt.Println("fields:", fields) // Output: fields: [Hello, World!]

	// FieldsFunc
	// FieldsFunc splits s around each instance of one or more consecutive Unicode code points satisfying f(r).
	fields = strings.FieldsFunc("  Hello,   World!  ", func(r rune) bool {
		return r == ' ' || r == ','
	})
	fmt.Println("fields:", fields) // Output: fields: [Hello World!]

	// FieldsFuncSeq
	// FieldsFuncSeq returns an iterator over substrings of s split around runs of Unicode code points satisfying f(c).
	iter := strings.FieldsFuncSeq("  Hello,   World!  ", func(r rune) bool {
		return r == ' ' || r == ','
	})
	for v := range iter {
		fmt.Println(v) // Output: Hello World!
	}

	// FieldsSeq
	// FieldsSeq returns an iterator over substrings of s split around runs of Unicode code points satisfying f(c).
	iter = strings.FieldsSeq("  Hello,   World!  ")
	for v := range iter {
		fmt.Println(v) // Output: Hello World!
	}
}

// String Case Functions
// Example of functions that manipulate the case of strings.
func StringCaseFunctions() {

	// ToLower
	// ToLower returns a copy of the string s with all Unicode code points mapped to their lower case.
	x := strings.ToLower("Hello, World!")
	fmt.Println("x:", x) // Output: x: hello, world!

	// ToLowerSpecial
	// ToLowerSpecial returns a copy of the string s with all Unicode code points mapped to their lower case.
	x = strings.ToLowerSpecial(unicode.TurkishCase, "Hello, World!")
	fmt.Println("x:", x) // Output: x: hello, world!

	// ToTitle
	// ToTitle returns a copy of the string s with all Unicode code points mapped to their title case.
	x = strings.ToTitle("hello world!")
	fmt.Println("x:", x) // Output: x: Hello World!

	// ToTitleSpecial
	// ToTitleSpecial returns a copy of the string s with all Unicode code points mapped to their title case.
	x = strings.ToTitleSpecial(unicode.TurkishCase, "hello world!")
	fmt.Println("x:", x) // Output: x: Hello World!

	// ToUpper
	// ToUpper returns a copy of the string s with all Unicode code points mapped to their upper case.
	x = strings.ToUpper("Hello, World!")
	fmt.Println("x:", x) // Output: x: HELLO, WORLD!

	// ToUpperSpecial
	// ToUpperSpecial returns a copy of the string s with all Unicode code points mapped to their upper case.
	x = strings.ToUpperSpecial(unicode.TurkishCase, "Hello, World!")
	fmt.Println("x:", x) // Output: x: HELLO, WORLD!
}

// String Trim Functions
// Example of functions that trim characters from strings.
func StringTrimFunctions() {

	// Trim
	// Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
	x := strings.Trim("  Hello, World!  ", " ")
	fmt.Println("x:", x) // Output: x: Hello, World! (Leading and trailing spaces removed)

	// TrimFunc
	// TrimFunc returns a slice of the string s with all leading and trailing Unicode code points satisfying f(c) removed.
	x = strings.TrimFunc("  Hello, World!  ", func(r rune) bool {
		return r == ' '
	})
	fmt.Println("x:", x) // Output: x: Hello, World! (Leading and trailing spaces removed)

	// TrimLeft
	// TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed.
	x = strings.TrimLeft("  Hello, World!  ", " ")
	fmt.Println("x:", x) // Output: x: Hello, World!  (Leading spaces removed)

	// TrimLeftFunc
	// TrimLeftFunc returns a slice of the string s with all leading Unicode code points satisfying f(c) removed.
	x = strings.TrimLeftFunc("  Hello, World!  ", func(r rune) bool {
		return r == ' '
	})
	fmt.Println("x:", x) // Output: x: Hello, World!  (Leading spaces removed)

	// TrimPrefix
	// TrimPrefix returns a slice of the string s with the provided leading prefix string removed.
	x = strings.TrimPrefix("Hello, World!", "Hello")
	fmt.Println("x:", x) // Output: x: , World! (Prefix 'Hello' removed)

	// TrimRight
	// TrimRight returns a slice of the string s with all trailing Unicode code points contained in cutset removed.
	x = strings.TrimRight("  Hello, World!  ", " ")
	fmt.Println("x:", x) // Output: x:   Hello, World! (Trailing spaces removed)

	// TrimRightFunc
	// TrimRightFunc returns a slice of the string s with all trailing Unicode code points satisfying f(c) removed.
	x = strings.TrimRightFunc("  Hello, World!  ", func(r rune) bool {
		return r == ' '
	})
	fmt.Println("x:", x) // Output: x:   Hello, World! (Trailing spaces removed)

	// TrimSpace
	// TrimSpace returns a slice of the string s with all leading and trailing white space removed.
	x = strings.TrimSpace("  Hello, World!  ")
	fmt.Println("x:", x) // Output: x: Hello, World! (Leading and trailing spaces removed)

	// TrimSuffix
	// TrimSuffix returns a slice of the string s with the provided trailing suffix string removed.
	x = strings.TrimSuffix("Hello, World!", "World!")
	fmt.Println("x:", x) // Output: x: Hello,  (Suffix 'World!' removed)
}

// String Other Functions
// Example of other functions in the strings package.
func StringOtherFunctions() {

	// Clone
	// Clone is used to create a copy of a string.
	x := "Hello, World!"
	y := strings.Clone(x)
	fmt.Println("y:", y) // Output: y: Hello, World!

	// Count
	// Count counts the number of non-overlapping instances of substr in s.
	n := strings.Count("Hello, World!", "o")
	fmt.Println("n:", n) // Output: n: 2

	// Join
	// Join concatenates the elements of a to create a single string.
	// The separator string sep is placed between elements in the resulting string.
	x = strings.Join([]string{"Hello", "World"}, ", ")
	fmt.Println("x:", x) // Output: x: Hello, World

	// Lines
	// Lines returns a slice of strings, each representing a line in s.
	lines := strings.Lines("Hello\nWorld!")
	fmt.Println("lines:", lines) // Output: lines: [Hello World!]

	// Map
	// Map returns a copy of the string s with all Unicode code points mapped by the mapping function.
	// The mapping function is applied to each Unicode code point in s.
	x = strings.Map(func(r rune) rune {
		if r == 'l' {
			return '1'
		}
		return r
	}, "Hello")
	fmt.Println("x:", x) // Output: x: He11o

	// Repeat
	// Repeat returns a new string consisting of count copies of the string s.
	x = strings.Repeat("Hello", 3)
	fmt.Println("x:", x) // Output: x: HelloHelloHello

	// ToValidUTF8
	// ToValidUTF8 returns a copy of the string s with invalid UTF-8 sequences replaced by the replacement character.
	x = strings.ToValidUTF8("Hello, \xFFWorld!", " ")
	fmt.Println("x:", x) // Output: x: Hello,  World! (Invalid UTF-8 replaced by space)
}

// String Types
// Example of types in the strings package.
func StringTypes() {

	// Builder
	// The Builder type is used to efficiently build strings.
	// It minimizes memory copying.
	// The zero value is ready to use.
	b := strings.Builder{}
	b.Grow(15) // Preallocate space for 100 bytes
	b.WriteString("Hello")
	b.WriteByte(' ')
	b.WriteString("World!")
	b.WriteRune('!')
	fmt.Println("b:", b.String()) // Output: b: Hello World!

	// Reader
	// The Reader type is used to read strings.
	// It implements the io.Reader interface.
	// It is used to read strings in a streaming fashion.
	r := strings.NewReader("Hello, World!")
	buf := make([]byte, 64)
	r.Read(buf)
	fmt.Println("buf:", string(buf)) // Output: buf: Hello, World!

	// Replacer
	// The Replacer type is used to replace strings in a string.
	// It is used to replace multiple strings in a single pass.
	// It is more efficient than using Replace multiple times.
	replacer := strings.NewReplacer("Hello", "Hi", "World", "Earth")
	x := replacer.Replace("Hello, World!")
	fmt.Println("x:", x) // Output: x: Hi, Earth! (Replaced 'Hello' with 'Hi' and 'World' with 'Earth')
}
