package main

import (
	"adventofcode/day_07/disc"
	"adventofcode/day_07/input"
	"fmt"
)

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}

// PartOne Solves the first part of day six of Advent of Code
func PartOne() string {
	// disc.PrintDisc(disc.ParseDisc(input.Raw), 0)
	disc := disc.ParseDisc(input.Raw)
	return disc.Name
}

// PartTwo Solves the second part of day six of Advent of Code
func PartTwo() int {
	d := disc.ParseDisc(input.Raw)
	return disc.FindNewWeight(&d)
}
