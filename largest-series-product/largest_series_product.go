//Package lsproduct returns the largest product for n digits in a series.
package lsproduct

import (
	"errors"
	"unicode"
)

const testVersion = 5

//LargestSeriesProduct calculates the largest product for n digits in a series.
func LargestSeriesProduct(s string, n int) (p int64, err error) {
	strln := int(len(s))
	switch {
	case strln < n || n < 0:
		return 0, errors.New("Invalid span length")
	case n == 0:
		return 1, nil
	default:
		var max int64
		p = 0
		for i := 0; i <= strln-n; i++ {
			if !unicode.IsDigit(rune(s[i])) {
				return 0, errors.New("Non-numeric digit found in series")
			}
			series := s[i : i+n]
			max = 1
			for _, val := range series {
				max *= int64(val - '0')
			}
			if max > p {
				p = max
			}
		}
	}
	return p, nil
}
