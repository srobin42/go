//Package sieve uses the Sieve of Eratosthenes to find all the primes from 2 up to a given number.
package sieve

const testVersion = 1

//Sieve finds all primes from 2 to i using the Sieve of Eratosthenes algorithm.
func Sieve(n int) (primes []int) {
	list := make([]int, 0, n-2)

	//create intial range of numbers to sieve
	for i := 2; i <= n; i++ {
		list = append(list, i)
	}
	for i := 0; i < len(list); i++ {
		num := list[i]
		if num*num > n {
			break
		}
		for j := 0; j < len(list); j++ {
			val := list[j]
			if num != 0 && val%num == 0 && num != val {
				list[j] = 0
			}
		}
	}
	for _, val := range list {
		if val != 0 {
			primes = append(primes, val)
		}
	}

	return primes
}
