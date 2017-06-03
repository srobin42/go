//Package acronym creates an acronym from a longer phrase, returns string.
package acronym

import (
	"strings"
)

const testVersion = 2

//Abbreviate creates an acronym from input string, returns string.
func Abbreviate(orig string) string {
	var out string

	//take care of special chars in original string
	orig = strings.Replace(orig, "-", " ", -1)
	orig = strings.Replace(orig, ",", "", -1)

	a := strings.Split(orig, " ")
	for _, word := range a {
		word = strings.Title(word)
		out += word[:1]
		//check for words with cap letters not in first pos
		include := false
		for i := 1; i < len(word); i++ {
			l := string(word[i])
			if include {
				if l == strings.ToUpper(l) {
					out += l
					include = !include
				}
			} else {
				if l != strings.ToUpper(l) {
					include = !include
				}
			}
		}
	}
	return out
}
