//Package luhn uses the Luhn formula to determine a valid number.
package luhn

import (
	"strings"
	"unicode"
)

const testVersion = 2

//Valid tests the validity of the input using the Luhn formula.
func Valid(card string) bool {
	card = strings.Replace(card, " ", "", -1)
	if len(card) < 2 {
		return false
	}

	sum := 0
	double := true
	if len(card)%2 != 0 {
		double = false
	}
	for _, val := range card {
		if unicode.IsDigit(val) {
			digit := int(val - '0')
			if double {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			double = !double
			sum += digit
		} else {
			return false
		}
	}
	/* submitted this, but return statement below is better
	if sum%10 == 0 {
		return true
	}
	return false
	*/
	return sum%10 == 0
}
