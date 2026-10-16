// Strings
// - Strings is a package that provides functions to manipulate strings.
// - It includes many utility functions for string operations.

package library

import (
	"fmt"
	"strings"
	"unicode"
)

//==================================================================================================
// Compare Functions
//==================================================================================================

func StringCompareFunctions() {

	// Compare
	// - Compare compares two strings lexicographically.
	// - It returns: 0 if a == b | -1 if a < b | +1 if a > b
	// - Output: cmp: -1

	cmp := strings.Compare("Hello", "World")
	fmt.Println("cmp:", cmp)

	// EqualFold
	// - EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under Unicode
	//   case-folding.
	// - It is case-insensitive.
	// - Output: eq: true

	eq := strings.EqualFold("Hello", "hello")
	fmt.Println("eq:", eq)
}

//==================================================================================================
// Contains Functions
//==================================================================================================

func StringContainsFunctions() {

	// Contains
	// - Contains reports whether substr is within s.
	// - Output: has: true

	has := strings.Contains("Hello, World!", "World")
	fmt.Println("has:", has)

	// ContainsAny
	// - ContainsAny reports whether any Unicode code points in chars are within s.
	// - The example below checks if any vowels are present in the string.
	// - Output: has: true

	has = strings.ContainsAny("Hello, World!", "aeiou")
	fmt.Println("has:", has)

	// ContainsFunc
	// - ContainsFunc reports whether any Unicode code points r within s satisfy f(r).
	// - The example below checks if any vowels are present in the string using a custom function.
	// - Output: has: true

	has = strings.ContainsFunc("Hello, World!", func(r rune) bool {
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	})
	fmt.Println("has:", has)

	// ContainsRune
	// - ContainsRune reports whether the Unicode code point r is within s.
	// - Output: has: true

	has = strings.ContainsRune("Hello, World!", 'H')
	fmt.Println("has:", has)

	// HasPrefix
	// - HasPrefix reports whether the string s begins with prefix.
	// - Output: has: true

	has = strings.HasPrefix("Hello, World!", "Hello")
	fmt.Println("has:", has)

	// HasSuffix
	// - HasSuffix reports whether the string s ends with suffix.
	// - Output: has: true

	has = strings.HasSuffix("Hello, World!", "World!")
	fmt.Println("has:", has)
}

//==================================================================================================
// Index Functions
//==================================================================================================

func StringIndexFunctions() {

	// Index
	// - Index returns the index of the first instance of substr in s, or -1 if substr is not
	//   present.
	// - Output: i: 7 (Found 'World' at index 7)

	i := strings.Index("Hello, World!", "World")
	fmt.Println("i:", i)

	// IndexAny
	// - IndexAny returns the index of the first instance of any Unicode code point in chars within
	//   s, or -1 if none are present.
	// - Output: i: 1 (Found 'e' at index 1)

	i = strings.IndexAny("Hello, World!", "aeiou")
	fmt.Println("i:", i)

	// IndexByte
	// - IndexByte returns the index of the first instance of c in s, or -1 if c is not present.
	// - Output: i: 7 (Found 'W' at index 7)

	i = strings.IndexByte("Hello, World!", 'W')
	fmt.Println("i:", i)

	// IndexFunc
	// - IndexFunc returns the index of the first Unicode code point satisfying f(c) in s, or -1 if
	//   none do.
	// - Output: i: 7 (Found 'W' at index 7)

	i = strings.IndexFunc("Hello, World!", func(r rune) bool {
		return r == 'W'
	})
	fmt.Println("i:", i)

	// IndexRune
	// - IndexRune returns the index of the first instance of r in s, or -1 if r is not present.
	// - Output: i: 7 (Found 'W' at index 7)

	i = strings.IndexRune("Hello, World!", 'W')
	fmt.Println("i:", i)

	// LastIndex
	// - LastIndex returns the index of the last instance of substr in s, or -1 if substr is not
	//   present.
	// - Output: i: 8 (Found last 'o' at index 8)

	i = strings.LastIndex("Hello, World!", "o")
	fmt.Println("i:", i)

	// LastIndexAny
	// - LastIndexAny returns the index of the last instance of any Unicode code point in chars
	//   within s, or -1 if none are present.
	// - Output: i: 8 (Found last 'o' at index 8)

	i = strings.LastIndexAny("Hello, World!", "aeiou")
	fmt.Println("i:", i)

	// LastIndexByte
	// - LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present.
	// - Output: i: 8 (Found last 'o' at index 8)

	i = strings.LastIndexByte("Hello, World!", 'o')
	fmt.Println("i:", i)

	// LastIndexFunc
	// - LastIndexFunc returns the index of the last Unicode code point satisfying f(c) in s, or -1
	//   if none do.
	// - Output: i: 8 (Found last 'o' at index 8)

	i = strings.LastIndexFunc("Hello, World!", func(r rune) bool {
		return r == 'o'
	})
	fmt.Println("i:", i)
}

//==================================================================================================
// Replace Functions
//==================================================================================================

func StringReplaceFunctions() {

	// Replace
	// - Replace returns a copy of the string s with the first n non-overlapping instances of old
	//   replaced by new.
	// - If n < 0, all instances are replaced.
	// - Output: x: He1lo (Only the first 'l' is replaced)

	x := strings.Replace("Hello", "l", "1", 1)
	fmt.Println("x:", x)

	// ReplaceAll
	// - ReplaceAll returns a copy of the string s with all non-overlapping instances of old
	//   replaced by new.
	// - It is an alias for Replace(s, old, new, -1).
	// - Output: x: He11o (All 'l' are replaced)

	x = strings.ReplaceAll("Hello", "l", "1")
	fmt.Println("x:", x)
}

//==================================================================================================
// Cut Functions
//==================================================================================================

func StringCutFunctions() {

	// Cut
	// - Cut slices s into two parts: the text before the first instance of sep and the text after.
	// - It returns the text before sep, the text after sep, and a boolean indicating if sep was
	//   found.
	// - Output: Hello | World! | true

	before, after, found := strings.Cut("Hello, World!", ", ")
	fmt.Println(before, after, found)

	// CutPreffix
	// - CutPrefix returns s without the provided leading prefix string.
	// - It returns the string after the prefix and a boolean indicating if the prefix was found.
	// - Output: World! | true

	after, found = strings.CutPrefix("Hello, World!", "Hello, ")
	fmt.Println(after, found)

	// CutSuffix
	// - CutSuffix returns s without the provided trailing suffix string.
	// - It returns the string before the suffix and a boolean indicating if the suffix was found.
	// - Output: Hello | true

	before, found = strings.CutSuffix("Hello, World!", ", World!")
	fmt.Println(before, "|", found)
}

//==================================================================================================
// Split Functions
//==================================================================================================

func StringSplitFunctions() {

	// Split
	// - Split slices s into all substrings separated by sep and returns a slice of the substrings
	//   between those separators.
	// - If sep is empty, Split splits after each UTF-8 sequence.
	// - If sep is not found, Split returns a slice of s.
	// - If s is empty, Split returns an empty slice.
	// - Output: parts: [a b c]

	parts := strings.Split("a,b,c", ",")
	fmt.Println("parts:", parts)

	// SplitAfter
	// - SplitAfter slices s into all substrings after each instance of sep and returns a slice of
	//   the substrings.
	// - Output: parts: [a, b, c] (The separator is included in the result)

	parts = strings.SplitAfter("a,b,c", ",")
	fmt.Println("parts:", parts)

	// SplitAfterN
	// - SplitAfterN slices s into all substrings after each instance of sep and returns a slice of
	//   the substrings.
	// - The result is limited to n substrings.
	// - Output: parts: [a, b,c] (The separator is included in the result)

	parts = strings.SplitAfterN("a,b,c", ",", 2)
	fmt.Println("parts:", parts)

	// SplitAfterSeq
	// - SplitAfterSeq slices s into all substrings after each instance of sep and returns a slice
	//   of the substrings.
	// - Output: a, b, c (The separator is included in the result)

	iter := strings.SplitAfterSeq("a,b,c", ",")
	for v := range iter {
		fmt.Println(v)
	}

	// SplitN
	// - SplitN slices s into all substrings separated by sep and returns a slice of the substrings.
	// - The result is limited to n substrings.
	// - Output: parts: [a b,c] (The separator is not included in the result)

	parts = strings.SplitN("a,b,c", ",", 2)
	fmt.Println("parts:", parts)

	// SplitSeq
	// - SplitSeq slices s into all substrings separated by sep and returns a slice of the
	//   substrings.
	// - Output: a b c (The separator is not included in the result)

	iter = strings.SplitSeq("a,b,c", ",")
	for v := range iter {
		fmt.Println(v)
	}
}

//==================================================================================================
// Fields Functions
//==================================================================================================

func StringFieldsFunctions() {

	// Fields
	// - Fields splits s around each instance of one or more consecutive white space characters.
	// - Output: fields: [Hello, World!]

	fields := strings.Fields("  Hello,   World!  ")
	fmt.Println("fields:", fields)

	// FieldsFunc
	// - FieldsFunc splits s around each instance of one or more consecutive Unicode code points
	//   satisfying f(r).
	// - Output: fields: [Hello World!]

	fields = strings.FieldsFunc("  Hello,   World!  ", func(r rune) bool {
		return r == ' ' || r == ','
	})
	fmt.Println("fields:", fields)

	// FieldsFuncSeq
	// - FieldsFuncSeq returns an iterator over substrings of s split around runs of Unicode code
	//   points satisfying f(c).
	// - Output: Hello World!

	iter := strings.FieldsFuncSeq("  Hello,   World!  ", func(r rune) bool {
		return r == ' ' || r == ','
	})
	for v := range iter {
		fmt.Println(v)
	}

	// FieldsSeq
	// - FieldsSeq returns an iterator over substrings of s split around runs of Unicode code points
	//   satisfying f(c).
	// - Output: Hello World!

	iter = strings.FieldsSeq("  Hello,   World!  ")
	for v := range iter {
		fmt.Println(v)
	}
}

//==================================================================================================
// Case Functions
//==================================================================================================

func StringCaseFunctions() {

	// ToLower
	// - ToLower returns a copy of the string s with all Unicode code points mapped to their lower
	//   case.
	// - Output: x: hello, world!

	x := strings.ToLower("Hello, World!")
	fmt.Println("x:", x)

	// ToLowerSpecial
	// - ToLowerSpecial returns a copy of the string s with all Unicode code points mapped to their
	//   lower case.
	// - Output: x: hello, world!

	x = strings.ToLowerSpecial(unicode.TurkishCase, "Hello, World!")
	fmt.Println("x:", x)

	// ToTitle
	// - ToTitle returns a copy of the string s with all Unicode code points mapped to their title
	//   case.
	// - Output: x: Hello World!

	x = strings.ToTitle("hello world!")
	fmt.Println("x:", x)

	// ToTitleSpecial
	// - ToTitleSpecial returns a copy of the string s with all Unicode code points mapped to their
	//   title case.
	// - Output: x: Hello World!

	x = strings.ToTitleSpecial(unicode.TurkishCase, "hello world!")
	fmt.Println("x:", x)

	// ToUpper
	// - ToUpper returns a copy of the string s with all Unicode code points mapped to their upper
	//   case.
	// - Output: x: HELLO, WORLD!

	x = strings.ToUpper("Hello, World!")
	fmt.Println("x:", x)

	// ToUpperSpecial
	// - ToUpperSpecial returns a copy of the string s with all Unicode code points mapped to their
	//   upper case.
	// - Output: x: HELLO, WORLD!

	x = strings.ToUpperSpecial(unicode.TurkishCase, "Hello, World!")
	fmt.Println("x:", x)
}

//==================================================================================================
// Trim Functions
//==================================================================================================

func StringTrimFunctions() {

	// Trim
	// - Trim returns a string with all leading and trailing Unicode code points contained in cutset
	//   removed.
	// - Output: x: Hello (Leading and trailing spaces removed)

	x := strings.Trim("  Hello  ", " ")
	fmt.Println("x:", x)

	// TrimFunc
	// - TrimFunc returns a string with all leading and trailing Unicode code points satisfying f(c)
	//   removed.
	// - Output: x: Hello (Leading and trailing spaces removed)

	x = strings.TrimFunc("  Hello  ", func(r rune) bool {
		return r == ' '
	})
	fmt.Println("x:", x)

	// TrimLeft
	// - TrimLeft returns a string with all leading Unicode code points contained in cutset removed.
	// - Output: x: Hello  (Leading spaces removed)

	x = strings.TrimLeft("  Hello  ", " ")
	fmt.Println("x:", x)

	// TrimLeftFunc
	// - TrimLeftFunc returns a string with all leading Unicode code points satisfying f(c) removed.
	// - Output: x: Hello  (Leading spaces removed)

	x = strings.TrimLeftFunc("  Hello  ", func(r rune) bool {
		return r == ' '
	})
	fmt.Println("x:", x)

	// TrimPrefix
	// - TrimPrefix returns a string with the provided leading prefix string removed.
	// - Output: x: llo (Prefix 'He' removed)

	x = strings.TrimPrefix("Hello", "He")
	fmt.Println("x:", x)

	// TrimRight
	// - TrimRight returns a string with all trailing Unicode code points contained in cutset
	//   removed.
	// - Output: x:   Hello (Trailing spaces removed)

	x = strings.TrimRight("  Hello  ", " ")
	fmt.Println("x:", x)

	// TrimRightFunc
	// - TrimRightFunc returns a string with all trailing Unicode code points satisfying f(c)
	//   removed.
	// - Output: x:   Hello (Trailing spaces removed)

	x = strings.TrimRightFunc("  Hello  ", func(r rune) bool {
		return r == ' '
	})
	fmt.Println("x:", x)

	// TrimSpace
	// - TrimSpace returns a string with all leading and trailing white space removed.
	// - Output: x: Hello (Leading and trailing spaces removed)

	x = strings.TrimSpace("  Hello  ")
	fmt.Println("x:", x)

	// TrimSuffix
	// - TrimSuffix returns a string with the provided trailing suffix string removed.
	// - Output: x: He (Suffix 'llo' removed)

	x = strings.TrimSuffix("Hello", "llo")
	fmt.Println("x:", x)
}

//==================================================================================================
// Other Functions
//==================================================================================================

func StringOtherFunctions() {

	// Clone
	// - Clone is used to create a copy of a string.
	// - Output: y: Hello, World!

	x := "Hello, World!"
	y := strings.Clone(x)
	fmt.Println("y:", y)

	// Count
	// - Count counts the number of non-overlapping instances of substr in s.
	// - Output: n: 2

	n := strings.Count("Hello, World!", "o")
	fmt.Println("n:", n)

	// Join
	// - Join concatenates the elements of a to create a single string.
	// - The separator string sep is placed between elements in the resulting string.
	// - Output: x: Hello, World

	x = strings.Join([]string{"Hello", "World"}, ", ")
	fmt.Println("x:", x)

	// Lines
	// - Lines returns an iterator over the newline-terminated lines in a string.
	// - Output: lines: Hello | World!
	iter := strings.Lines("Hello\nWorld!")
	for line := range iter {
		fmt.Println(line)
	}

	// Map
	// - Map returns a copy of the string s with all Unicode code points mapped by the mapping
	//   function.
	// - The mapping function is applied to each Unicode code point in s.
	// - Output: x: He11o

	x = strings.Map(func(r rune) rune {
		if r == 'l' {
			return '1'
		}
		return r
	}, "Hello")
	fmt.Println("x:", x)

	// Repeat
	// - Repeat returns a new string consisting of count copies of the string s.
	// - Output: x: HelloHelloHello

	x = strings.Repeat("Hello", 3)
	fmt.Println("x:", x)

	// ToValidUTF8
	// - ToValidUTF8 returns a copy of the string s with invalid UTF-8 sequences replaced by the
	//   replacement character.
	// - Output: x: Hello,  World! (Invalid UTF-8 replaced by space)

	x = strings.ToValidUTF8("Hello, \xFFWorld!", " ")
	fmt.Println("x:", x)
}

//==================================================================================================
// String Types
//==================================================================================================

func StringTypes() {

	// Builder
	// - The Builder type is used to efficiently build strings.
	// - It minimizes memory copying.
	// - The zero value is ready to use.
	// - Output: Hello World!

	b := strings.Builder{}
	b.Grow(15) // Preallocate space for 15 bytes
	b.WriteString("Hello")
	b.WriteByte(' ')
	b.WriteString("World")
	b.WriteRune('!')
	fmt.Println(b.String())

	// Reader
	// - The Reader type is used to read strings.
	// - It implements the io.Reader interface.
	// - It is used to read strings in a streaming fashion.
	// - Output: buf: Hello, World!

	r := strings.NewReader("Hello, World!")
	buf := make([]byte, 13)
	r.Read(buf)
	fmt.Println("buf:", string(buf))

	// Replacer
	// - The Replacer type is used to replace strings in a string.
	// - It is used to replace multiple strings in a single pass.
	// - It is more efficient than using Replace multiple times.
	// - Output: x: Hi, Earth! (Replaced 'Hello' with 'Hi' and 'World' with 'Earth')

	replacer := strings.NewReplacer("Hello", "Hi", "World", "Earth")
	x := replacer.Replace("Hello, World!")
	fmt.Println("x:", x)
}
