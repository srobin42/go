//Package etl transforms legacy data into current format.
package etl

import "strings"

const testVersion = 1

//Transform takes legacy data and transforms into current format.
func Transform(m map[int][]string) map[string]int {
	r := make(map[string]int)
	for score, val := range m {
		for _, l := range val {
			r[strings.ToLower(l)] = score
		}
	}
	return r
}
