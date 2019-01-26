package main

import (
	"adventofcode/2017/day_03/one"
	"adventofcode/2017/day_03/two"
	"fmt"
)

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}

const input int = 325489

// PartOne Solves the first part of day three of Advent of Code
func PartOne() int {
	return one.LateralDistance(input) + one.RotationalDistance(input)
}

// PartTwo Solves the second part of day three of Advent of Code
func PartTwo() int {
	return two.CountTo(input)
}
