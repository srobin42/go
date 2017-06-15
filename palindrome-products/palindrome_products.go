//Package palindrome detects palindrome products in a given range.
package palindrome

import (
	"errors"
)

const testVersion = 1

//Product structure holds the palindrome product and factor ints.
type Product struct {
	Product        int
	Factorizations [][2]int
}

//Products returns the min and max palindrome product types in a given range.
func Products(fmin, fmax int) (pmin, pmax Product, ok error) {
	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax")
	}
	min := fmin

	for min <= fmax {
		for i := fmin; i <= fmax; i++ {
			prod := min * i
			if isPalindrome(prod) {
				if pmin.Product == 0 {
					pmin.Product = prod
					pmin.Factorizations = [][2]int{{min, i}}
				}
				if prod >= pmax.Product {
					pmax.Product = prod
				}
			}
		}
		min++
	}

	//check for no palindromes
	if pmin.Product == 0 {
		ok = errors.New("no palindromes")
	} else {
		//get factorizations for max product
		pmax.Factorizations = maxFactors(pmax.Product, fmin, fmax)
	}

	return pmin, pmax, ok
}

func isPalindrome(i int) bool {
	if i%10 == 0 {
		return false
	}
	test := i
	rev := 0
	for test > 0 {
		rev = rev*10 + test%10
		test = test / 10
	}
	return rev == i
}

func maxFactors(i, min, max int) (factors [][2]int) {
	//calculate factorizations
	for j := min; j < max; j++ {
		if i%j == 0 {
			if i/j <= max && j <= i/j {
				factors = append(factors, [2]int{j, i / j})
			}
		}
	}

	return factors
}
