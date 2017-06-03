//Package grains calculates number of grains on each square of chessboard and total grains used. Focused on code readability.
package grains

import (
	"errors"
)

const testVersion = 1

//Square returns the number of grains on a particular square.
func Square(i int) (t uint64, err error) {
	if i < 1 || i > 64 {
		err = errors.New("Number out of bounds")
		return 0, err
	}

	/* a much faster if hack-ish way to calc the square:
	shifts the bits by one until proper number returned.
	This works due to the nature of this problem: 1,2,4,8,16, etc.

	return 1 << uint(n-1), nil
	*/

	return sqGrains(uint64(i)), nil
}

//Total returns the total number of grains used.
func Total() uint64 {
	var t uint64
	for i := 64; i > 0; i-- {
		t += sqGrains(uint64(i))
	}

	/* better way to calc the total:
	for i:= 1; i <65; i++
		t *= 2
	}
	return t - 1
	*/
	return t
}

//sqGrains returns number of grains used on a certain square.
func sqGrains(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	return sqGrains(n-1) * 2
}
