// Package leap takes a year as input and calculates if that year
// is a leap year and returns a boolean.
package leap

const testVersion = 3

// IsLeapYear takes a year as input in integer format and calculates
// if the year is a leap year or not and returns a boolean.
func IsLeapYear(year int) bool {
	switch {
	case year%4 != 0:
		return false
	case year%100 == 0 && year%400 == 0:
		return true
	case year%100 == 0:
		return false
	default:
		return true
	}
}
