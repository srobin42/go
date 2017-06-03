//Package letter counts the frequency of letters in a string.
package letter

const testVersion = 1

//FreqMap holds the letter and frequency count.
type FreqMap map[rune]int

//Frequency counts the frequency of letters in a string.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency counts the frequency of letters in multiple strings concurrently.
func ConcurrentFrequency(s []string) FreqMap {
	m := FreqMap{}
	c := make(chan FreqMap)

	for _, v := range s {
		go func(verse string) {
			c <- Frequency(verse)
		}(v)
	}
	for i := 0; i < len(s); i++ {
		for key, val := range <-c {
			m[key] += val
		}
	}
	return m
}
