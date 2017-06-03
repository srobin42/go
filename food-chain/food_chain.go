//Package foodchain assembles a folk song verse by verse.
package foodchain

import (
	"strings"
)

const testVersion = 3

var animals = []string{
	"",
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

var secondary = []string{
	"",
	"I don't know why she swallowed the fly. Perhaps she'll die.",
	"wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
	"She's dead, of course!",
}

//Verse returns a specific verse.
func Verse(i int) (v string) {
	v = "I know an old lady who swallowed a " + animals[i] + ".\n"
	if i == 1 || i == 8 {
		v += secondary[i]
	} else {
		if i > 2 {
			v += secondary[i] + "\n"
			for j := i; j >= 4; j-- {
				v += "She swallowed the " + animals[j] + " to catch the " + animals[j-1] + ".\n"
			}
			v += "She swallowed the " + animals[3] + " to catch the " + animals[2] + " that " + secondary[2] + "\n"
		} else {
			v += "It " + secondary[2] + "\n"
		}
		v += "She swallowed the " + animals[2] + " to catch the " + animals[1] + ".\n"
		v += secondary[1]
	}
	return v
}

//Verses returns all verses between starting and ending points passsed in.
func Verses(start int, end int) (v string) {
	for ; start <= end; start++ {
		v += Verse(start) + "\n\n"
	}
	return strings.TrimSpace(v)
}

//Song returns the entire song
func Song() string {
	return strings.TrimSpace(Verses(1, 8))
}
