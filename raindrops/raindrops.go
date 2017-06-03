/*Package raindrops converts input int to string and returns a
string containing certain words depending on the input's factors:
  factors 3 - 'Pling', 5 - 'Plang', 7 - 'Plong' */
package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 3

// Convert takes an int, converts into string and returns a string
func Convert(i int) string {
	var buffer bytes.Buffer

	if i%3 == 0 {
		buffer.WriteString("Pling")
	}
	if i%5 == 0 {
		buffer.WriteString("Plang")
	}
	if i%7 == 0 {
		buffer.WriteString("Plong")
	}

	if buffer.Len() == 0 {
		return strconv.Itoa(i)
	}

	return buffer.String()
}
