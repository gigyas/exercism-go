package hamming

import (
	"fmt"
)

const testVersion = 5

// Distance computes the hamming distance of two dns strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("Strands different length")
	}

	dist := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist, nil
}
