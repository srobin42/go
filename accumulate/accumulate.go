//Package accumulate performs specified function on input string array, outputs string array.
package accumulate

const testVersion = 1

//Accumulate performs function on input string array and returns results in string array.
func Accumulate(s []string, t func(string) string) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = t(v)
	}
	return r
}
