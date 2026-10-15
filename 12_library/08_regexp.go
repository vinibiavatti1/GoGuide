// Regexp
// - Regexp implements regular expression search and text manipulation.
// - It uses the RE2 syntax, which guarantees linear execution time relative to the text size.

package library

import (
	"fmt"
	"regexp"
)

//==================================================================================================
// Core Functions
//==================================================================================================

func RegexpCoreFunctions() {

	// MatchString
	// - MatchString reports whether the string contains any match of the regular expression
	//   pattern.
	// - It is the simplest and most direct way, without the need to explicitly compile.
	// - Output: match: true

	match, _ := regexp.MatchString(`\d+`, "123")
	fmt.Println("match:", match)

	// Compile
	// - Compile parses a regular expression and, if successful, returns a *Regexp object.
	// - It should be used when you need to reuse the expression or use advanced methods.
	// - Output: re: \d+

	re, _ := regexp.Compile(`\d+`)
	fmt.Println("re:", re)

	// MustCompile
	// - MustCompile is like Compile but panics if the expression cannot be parsed.
	// - It is ideal for global initialization or when you are certain the regex is valid.
	// - Output: re: [A-Z]+

	re = regexp.MustCompile(`[A-Z]+`)
	fmt.Println("re:", re)
}

//==================================================================================================
// Find Functions
//==================================================================================================

func RegexpFindFunctions() {
	re := regexp.MustCompile(`[A-Z]+`)

	// FindString
	// - FindString returns a string holding the text of the first match in the target text.
	// - If there is no match, the return value is an empty string.
	// - Output: match: ABC

	match := re.FindString("ABC DEF GHI")
	fmt.Println("match:", match)

	// FindAllString
	// - FindAllString returns a slice of all successive matches of the expression.
	// - The second argument defines the limit of results (-1 returns all matches).
	// - Output: matches: [ABC DEF GHI]

	matches := re.FindAllString("ABC DEF GHI", -1)
	fmt.Println("matches:", matches)

	// FindStringIndex
	// - FindStringIndex returns a two-element slice of integers defining the location of the first
	//   match (inclusive/exclusive).
	// - Output: loc: [0 3] (Indicates that 'ABC' is located between indices 0 and 3)

	loc := re.FindStringIndex("ABC DEF GHI")
	fmt.Println("loc:", loc)
}

//==================================================================================================
// Group Functions
//==================================================================================================

func RegexpGroupFunctions() {

	// FindStringSubmatch
	// - Returns a slice containing the text of the leftmost match followed by the matches of its
	//   subexpressions.
	// - Output: submatch: [2000-01-15 2000 01 15]

	re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	submatch := re.FindStringSubmatch("2000-01-15")
	fmt.Println("submatch:", submatch)

	// FindStringSubmatch (Named Groups)
	// - Go allows capturing named groups using the (?P<name>expr) syntax.
	// - We use SubexpNames() to extract and map these captured values.
	// - Output: year: 2000 | month: 01 | day: 15

	reNamed := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)
	submatchNamed := reNamed.FindStringSubmatch("2000-01-15")
	names := reNamed.SubexpNames()
	result := make(map[string]string)
	for i, name := range names {
		if i != 0 && name != "" {
			result[name] = submatchNamed[i]
		}
	}
	fmt.Printf("year: %s | month: %s | day: %s\n", result["year"], result["month"], result["day"])
}

//==================================================================================================
// Replace Functions
//==================================================================================================

func RegexpReplaceFunctions() {
	re := regexp.MustCompile(`\d+`)

	// ReplaceAllString
	// - ReplaceAllString returns a copy of the source string, replacing matches of the Regexp with
	//   the replacement string.
	// - Output: res: ABC-X-DEF-X

	res := re.ReplaceAllString("ABC-123-DEF-456", "X")
	fmt.Println("res:", res)

	// ReplaceAllStringFunc
	// - ReplaceAllStringFunc returns a copy of the source string where all matches are replaced via
	//   a custom function.
	// - Output: resFunc: ABC-[123]-DEF-[456]

	resFunc := re.ReplaceAllStringFunc("ABC-123-DEF-456", func(match string) string {
		return "[" + match + "]"
	})
	fmt.Println("resFunc:", resFunc)
}

//==================================================================================================
// Regexp Flags
//==================================================================================================

func RegexpFlags() {

	// Case Insensitive (?i)
	// - The (?i) flag makes the pattern matching case-insensitive.
	// - Output: matchCase: true

	reCase := regexp.MustCompile(`(?i)[a-z]+`)
	matchCase := reCase.MatchString("ABC")
	fmt.Println("matchCase:", matchCase)

	// Multiline (?m)
	// - The (?m) flag causes ^ and $ to match the beginning and end of each line, not just the
	//   entire text.
	// - Output: linesMatches: [ABC DEF GHI]

	reMulti := regexp.MustCompile(`(?m)^[A-Z]+$`)
	linesMatches := reMulti.FindAllString("ABC\nDEF\nGHI", -1)
	fmt.Println("linesMatches:", linesMatches)

	// Dot All (?s)
	// - The (?s) flag causes the dot (.) character to match newline characters (\n) as well.
	// - Output: matchDot: true

	reDot := regexp.MustCompile(`(?s).`)
	matchDot := reDot.MatchString("\n")
	fmt.Println("matchDot:", matchDot)
}
