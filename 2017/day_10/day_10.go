package main

import (
	"fmt"
)

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}

// PartOne Solves the first part of day X of Advent of Code
func PartOne() int {
	cycle := newCycle()
	lengths := parseLengths(raw)
	for i := range lengths {
		cycle.run(lengths[i])
	}
	return cycle.checksum()
}

// PartTwo Solves the second part of day X of Advent of Code
func PartTwo() string {
	cycle := newCycle()
	lengths := parseASCIILengths(raw)
	for n := 0; n < 64; n++ {
		for i := range lengths {
			cycle.run(lengths[i])
		}
	}
	return cycle.denseHash()
}
