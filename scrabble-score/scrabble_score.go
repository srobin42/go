//Package scrabble calculates word score.
package scrabble

import (
	"strings"
)

const testVersion = 5

//reduced values by one to reduce map size
var scores = map[rune]int{
	'b': 2,
	'c': 2,
	'd': 1,
	'f': 3,
	'g': 1,
	'h': 3,
	'j': 7,
	'k': 4,
	'm': 2,
	'p': 2,
	'q': 9,
	'v': 3,
	'w': 3,
	'x': 7,
	'y': 3,
	'z': 9,
}

//Score calculates scrabble score for a given word
func Score(w string) (s int) {
	w = strings.ToLower(w)
	for _, val := range w {
		s += scores[val]
		//add one to account for letters not in map
		s++
	}
	return s
}
