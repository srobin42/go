//Package bob takes an input string and figures out the correct response, outputs string.
package bob

import (
	"strings"
	"unicode"
)

const testVersion = 3

//Hey deciphers input string and crafts appropriate response, returns string.
func Hey(quest string) string {
	var ans string
	quest = strings.TrimSpace(quest)

	if len(quest) == 0 {
		ans = "Fine. Be that way!"
	} else {
		for _, val := range quest {
			if unicode.IsUpper(val) {
				ans = "Whoa, chill out!"
			} else {
				if !unicode.IsPunct(val) && !unicode.IsSpace(val) && !unicode.IsDigit(val) && !unicode.IsSymbol(val) {
					ans = ""
					break
				}
			}
		}

		if len(ans) == 0 {
			if strings.Contains(quest[len(quest)-2:], "?") {
				ans = "Sure."
			} else {
				ans = "Whatever."
			}
		}
	}

	return ans
}
