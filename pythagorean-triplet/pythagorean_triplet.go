//Package pythagorean returns Pythagorean triplets in a certain range
//or triplets where the perimeter is equal to sum of sides.
package pythagorean

import (
	"sort"
)

const testVersion = 1

//Triplet stores the side lengths of the triangle
type Triplet [3]int

//Range returns a list of all Pythagorean triplets with sides in the
//range min to max inclusive.
func Range(min, max int) (t []Triplet) {
	trips := genPPT()
	for _, val := range trips {
		if val[2] <= max {
			if val[0] > min {
				t = append(t, val)
			}
		}
	}
	return t
}

//Sum returns a list of all Pythagorean triplets where the sum a+b+c
//(the perimeter) is equal to p.
func Sum(p int) (t []Triplet) {
	var sol [][]int
	var prev int
	//get basic triplet sets
	basic := genPPT()
	for _, val := range basic {
		//check to see if a multiple of a basic set satisfies the requirement
		for m := 2; m < 200; m++ {
			if m*(val[0]+val[1]+val[2]) == p {
				sol = append(sol, []int{val[0] * m, val[1] * m, val[2] * m})
				break
			}
		}
	}
	//reverse the multiple sets and remove duplicates
	for i := len(sol) - 1; i >= 0; i-- {
		if sol[i][0] != prev {
			t = append(t, Triplet{sol[i][0], sol[i][1], sol[i][2]})
		}
		prev = sol[i][0]
	}
	return t
}

//generates primitive and (some) non primitive pythagorean triplets up to a limit
func genPPT() (t []Triplet) {
	var m, n int
	sides := make([]int, 3)

	for n = 1; n < 5; n++ {
		for m = n + 1; m < 6; m++ {
			sides[0] = (m * m) - (n * n)
			sides[1] = 2 * m * n
			sides[2] = (m * m) + (n * n)
			sort.Ints(sides)
			t = append(t, Triplet{sides[0], sides[1], sides[2]})
		}
	}

	return t
}
