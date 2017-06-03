//Package isogram identifies whether input is comprised of non-repeating characters
package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

//IsIsogram determines whether input is or is not an isogram.
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	for _, val := range s {
		if unicode.IsLetter(val) {
			if strings.Count(s, string(val)) > 1 {
				return false
			}
		}
	}
	return true
}
