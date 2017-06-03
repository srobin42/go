//Package twelve returns the lyrics to "12 Days of Christmas" song, output string.
package twelve

import (
	"bytes"
)

const testVersion = 1

//song map holds lyrics
var song = map[int]string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

//dName holds proper day possessive
var dName = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

//Verse returns a verse from song starting with day passed in going backwards, output string.
func Verse(v int) string {

	words := addVerses(v)
	r := "On the " + dName[v] + " day of Christmas my true love gave to me, " + words
	return r
}

//Song returns all twelve verses of song starting from the beginning, output string.
func Song() string {
	var buffer bytes.Buffer

	for i := 1; i <= 12; i++ {
		buffer.WriteString(Verse(i))
		buffer.WriteString("\n")
	}
	return buffer.String()
}

//addVerses concats all of the verses from day number passed in backwards, output strings.
func addVerses(i int) string {
	var buffer bytes.Buffer

	//compose song backwards from verse input to first verse
	num := i
	for num > 1 {
		buffer.WriteString(song[num])
		buffer.WriteString(", ")
		num--
	}
	//put and before first verse unless only verse one requested
	if i != 1 {
		buffer.WriteString("and ")
	}
	buffer.WriteString(song[1])

	buffer.WriteString(".")
	return buffer.String()
}
