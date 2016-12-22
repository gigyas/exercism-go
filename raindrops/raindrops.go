package raindrops

import (
	"strconv"
)

const testVersion = 2

// Convert converts the thing
func Convert(f int) string {
	var ret string

	if f%3 == 0 {
		ret += "Pling"
	}
	if f%5 == 0 {
		ret += "Plang"
	}
	if f%7 == 0 {
		ret += "Plong"
	}
	if len(ret) == 0 {
		ret = strconv.Itoa(f)
	}

	return ret
}
