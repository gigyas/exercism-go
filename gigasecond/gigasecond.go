// Package gigasecond implements the solution to http://exercism.io/exercises/go/gigasecond/readme
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4 // find the value in gigasecond_test.go

// AddGigasecond calculates the time 10^9 secnonds from the provided time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(1000000000) * time.Second)
}
