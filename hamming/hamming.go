// Package hamming counts variances in strings and returns the number of
// variance as an int, and an error message if strings are not the same.
package hamming

import (
	"errors"
)

const testVersion = 6

// Distance finds the variances between two strings, returning the number
// of differences and an error if the two strings do not match.
func Distance(a, b string) (int, error) {
	var errCnt int

	switch {
	case a == b:
		return 0, nil
	case len(a) != len(b):
		return 0, errors.New("String length not equal")
	default:
		//iterate over strings byte by byte and compare
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				errCnt++
			}
		}
		return errCnt, nil
	}
}
