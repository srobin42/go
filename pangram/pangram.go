//Package pangram takes an input string, parses it for every letter in the alphabet, returns a bool.
package pangram

import (
	"strings"
)

const testVersion = 1

//IsPangram takes the input string, returns true if every letter in alphabet is used in the string.
func IsPangram(orig string) bool {
	//create string of alphabet letters
	ab := "abcdefghijklmnopqrstuvwxyz"
	orig = strings.ToLower(orig)
	for _, val := range ab {
		if !strings.ContainsRune(orig, val) {
			return false
		}
	}
	return true
}
