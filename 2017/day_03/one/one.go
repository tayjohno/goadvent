package one

import "adventofcode/imath"

// LateralDistance returns which ring out from the center this number falls on.
func LateralDistance(input int) int {
	return ringNumber(input)
}

// RotationalDistance returns how far from an axis this falls
func RotationalDistance(input int) int {
	if input == 1 {
		return 0
	}
	n := ringNumber(input)
	e := ringEnd(n)
	i := e - input
	return imath.Abs(i%(n*2) - n)
}

// ringEnd returns the last number for a given ring
func ringEnd(r int) int {
	return 1 + 4*(r*r+r)
}

// ringNumber returns the ring that a number is found on
func ringNumber(input int) (d int) {
	for {
		if input <= ringEnd(d) {
			return
		}
		d++
	}
}
