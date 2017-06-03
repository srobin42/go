//Package queenattack determines whether two Chess queens can attack each other.
package queenattack

import (
	"errors"
	"strconv"
)

const testVersion = 2

//CanQueenAttack returns bool
func CanQueenAttack(w, b string) (attack bool, err error) {
	var wCol, bCol, wRow, bRow int
	//plot positions and attack vectors
	wCol, wRow, err = plotPositions(w)
	if err != nil {
		return false, err
	}
	bCol, bRow, err = plotPositions(b)
	if err != nil {
		return false, err
	}

	switch {
	case w == b:
		//queens occupy same square, no attack
		err = errors.New("Queens share same square")
	case wRow < 1 || bRow < 1 || wCol < 0 || bCol < 0:
		err = errors.New("One or more pieces off the board 1")
	case wRow > 8 || bRow > 8 || wCol > 7 || bCol > 7:
		err = errors.New("One or more pieces off the board 2")
	case wCol == bCol || wRow == bRow:
		//queens share row or column, can attack
		attack = true
	case (bRow-wRow) == (bCol-wCol) || (wRow-bCol) == (bRow-wCol):
		//queens on diagonal, can attack
		attack = true
	}

	return attack, err
}

func plotPositions(p string) (fi, si int, err error) {
	if len(p) < 2 {
		err = errors.New("Invalid entry")
	} else {
		f, s := p[0], p[1:]
		si, err = strconv.Atoi(s)
		fi = int(f) - int('a')
	}
	return fi, si, err
}
