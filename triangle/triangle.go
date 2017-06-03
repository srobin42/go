//Package triangle determines if points input create a triangle, if so returns type or NaT in output string.
package triangle

import (
	"math"
)

const testVersion = 3

//Kind string for triangle classification
type Kind string

//Triangle type const
const (
	Equ = Kind("Equ")
	Iso = Kind("Iso")
	Sca = Kind("Sca")
	NaT = Kind("NaT")
)

//KindFromSides takes 3 values and determines if they create a triangle, and what type if so, returns string.
func KindFromSides(a, b, c float64) Kind {
	var t Kind
	ab := a + b
	ac := a + c
	bc := b + c

	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		//Not a triangle, all sides must be a number
		t = "NaT"
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		//Not a triangle, no infinite sides
		t = "NaT"
	case a <= 0 || b <= 0 || c <= 0:
		//Not a triangle, all sides must have positive length
		t = "NaT"
	case ab < c || ac < b || bc < a:
		//Sum of two sides must be => than third side
		t = "NaT"
	case a == b && b == c:
		//Equilateral triangle
		t = "Equ"
	case a == b || b == c || a == c:
		//Isosceles triangle
		t = "Iso"
	default:
		//Scalene triangle
		t = "Sca"
	}
	return t
}
