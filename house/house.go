//Package house assembles the nursery rhyme 'This is the House that Jack Built' line by line, output string.
package house

import "bytes"
import "strings"

const testVersion = 1

var words string

var phrase = []string{
	"the house that Jack built.",
	"the malt",
	"the rat",
	"the cat",
	"the dog",
	"the cow with the crumpled horn",
	"the maiden all forlorn",
	"the man all tattered and torn",
	"the priest all shaven and shorn",
	"the rooster that crowed in the morn",
	"the farmer sowing his corn",
	"the horse and the hound and the horn",
}

var action = []string{
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

//Verse composes the input verse of the nursery rhyme, returns string.
func Verse(v int) string {
	return composeVerses(v, []string{})
}

//Song composes the entire nursery rhyme, returns string.
func Song() string {
	var buffer bytes.Buffer
	for i := 1; i <= 12; i++ {
		buffer.WriteString(Verse(i))
		buffer.WriteString("\n\n")
	}

	return strings.TrimSpace(buffer.String())
}

func composeVerses(v int, accum []string) string {
	if v == 0 {
		return strings.TrimSpace(strings.Join(accum, "\n"))
	}

	if len(accum) == 0 {
		accum = append(accum, "This is "+phrase[v-1])
	} else {
		accum = append(accum, "that "+action[v-1]+" "+phrase[v-1])

	}
	return composeVerses(v-1, accum)
}
