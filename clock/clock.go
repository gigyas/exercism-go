package clock

import (
	"fmt"
)

const testVersion = 4

// Clock defines a 24-hour clock
type Clock struct {
	hour, min int
}

// New creates a new 24-hour clock
func New(hour, minute int) Clock {
	newMin := minute % 60      //Clamp mins to [0,60)
	minuteHours := minute / 60 //How many more than 1 hour in minutes?
	if newMin < 0 {
		// If we have negative minutes left, add an hour
		newMin += 60
		minuteHours-- // and take one hour off
	}

	newHour := (hour + minuteHours) % 24
	if newHour < 0 {
		newHour += (-1 * newHour % 24) * 24
		newHour = newHour % 24
	}

	return Clock{
		hour: newHour,
		min:  newMin,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

// Add returns a new Clock with the sepcified number of minutes added
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.min+minutes)
}
