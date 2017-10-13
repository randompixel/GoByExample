package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Get the current time
	now := time.Now()
	p(now)

	// Create a time
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Extractions
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Questions
	p(then.Weekday())   // Tuesday
	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	// Calculations
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))  // advance time
	p(then.Add(-diff)) // subtract time

}
