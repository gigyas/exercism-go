// Package leap implements the solution to http://exercism.io/exercises/go/leap/readme
package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear returns true if the provided year is a leap year
func IsLeapYear(year int) bool {
	leap := true

	if year%4 != 0 {
		leap = false
	} else {
		if year%100 == 0 && year%400 != 0 {
			leap = false
		}
	}

	return leap
}
