//Package series returns n length strings from a series of digits.
package series

const testVersion = 2

//All returns all n length substrings from a series of digits.
func All(n int, s string) []string {
	subs := []string{}
	for i := 0; n <= len(s); i++ {
		subs = append(subs, s[i:n])
		n++
	}
	return subs
}

//UnsafeFirst returns the first n length substring of series.
func UnsafeFirst(n int, s string) string {
	var subs string
	if n <= len(s) {
		subs = s[0:n]
	}
	return subs
}

//First returns false if n > len of string passed in, else returns true, first n length substring.
func First(n int, s string) (first string, ok bool) {
	if n <= len(s) {
		first = s[0:n]
		ok = true
	}
	return first, ok

}
