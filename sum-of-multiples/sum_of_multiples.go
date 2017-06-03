//Package summultiples sums the multiples of input up to a limit.
package summultiples

const testVersion = 1

//SumMultiples sums the multiples of input up to a limit.
func SumMultiples(lim int, div ...int) (sum int) {
	m := make(map[int]struct{})
	for _, val := range div {
		if val < lim {
			m[val] = struct{}{}
		}
		for i := 1; (i * val) < lim; i++ {
			m[i*val] = struct{}{}
		}
	}

	for key := range m {
		sum += key
	}
	return sum
}
