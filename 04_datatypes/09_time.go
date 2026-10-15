// Time
// - The time package provides functionality for measuring and displaying time.
// - The core type is time.Time, which represents an instant in time with nanosecond precision.
// - Go uses a unique layout-based approach for formatting and parsing dates using a reference time.

package datatypes

import (
	"fmt"
	"time"
)

//==================================================================================================
// Declaring Time
//==================================================================================================

func DeclaringTime() {

	// Getting the Current Time
	// - We can use the "time.Now()" function to get the current local time.
	now := time.Now()
	fmt.Println("Now:", now)

	// Creating a Specific Date
	// - We can construct a date using "time.Date(year, month, day, hour, min, sec, nsec, loc)".
	// - Note: Months use typed constants like time.January, time.February, etc.
	// - Output: 2026-07-17 13:30:00 +0000 UTC
	specificDate := time.Date(2026, time.July, 17, 13, 30, 0, 0, time.UTC)
	fmt.Println("Specific Date:", specificDate)

	// Extracting Components from Time
	// - We can extract individual components using specific methods.
	// - Output: 2026 July 17 Friday
	t := time.Now()
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Weekday())
}

//==================================================================================================
// Formatting & Parsing
//==================================================================================================

func FormattingAndParsingTime() {
	t := time.Date(2026, time.July, 17, 13, 30, 0, 0, time.UTC)

	// Formatting Time (Converting Time to String)
	// - Go does NOT use traditional yyyy-MM-dd layout tokens, instead, it uses a specific reference
	//   time: "Mon Jan 2 15:04:05 MST 2006" (Easy to remember as: 1 2 3 4 5 6 7).
	// - We use the ".Format()" method passing our desired layout using the reference values.
	// - Output: 2026-07-17
	fmt.Println(t.Format("2006-01-02"))

	// Parsing Time (Converting String to Time)
	// - We use "time.Parse(layout, value)" to convert a string into a time.Time object.
	// - Output: 2026-07-17 13:30:00 +0000 UTC
	dateStr := "2026-07-17 13:30:00"
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", dateStr)
	fmt.Println("Parsed Time:", parsedTime)
}

//==================================================================================================
// Time Operations
//==================================================================================================

func TimeOperations() {
	t := time.Date(2026, time.July, 17, 13, 30, 0, 0, time.UTC)

	// Adding Durations
	// - We can add hours, minutes, seconds using predefined time.Duration constants.
	// - Output: 2026-07-18 13:30:00 +0000 UTC
	tomorrow := t.Add(24 * time.Hour)
	fmt.Println("Tomorrow:", tomorrow)

	// Adding Years, Months and Days
	// - For calendar-based additions, use ".AddDate(years, months, days)".
	// - Output: 2026-08-17 13:30:00 +0000 UTC
	nextMonth := t.AddDate(0, 1, 0)
	fmt.Println("Next Month:", nextMonth)

	// Substracting Time (Calculating Difference)
	// - The ".Sub()" method calculates the duration between two times.
	// - Output: Hours difference: 168
	pastDate := time.Date(2026, time.July, 10, 13, 30, 0, 0, time.UTC)
	duration := t.Sub(pastDate)
	fmt.Println("Hours difference:", duration.Hours())

	// Comparing Timelines
	// - We can use Before(), After(), and Equal() to compare time instances.
	// - Output: true false false
	fmt.Println(pastDate.Before(t), pastDate.After(t), pastDate.Equal(t))
}
