package main

import (
	"fmt"
)

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}

// PartOne Solves the first part of day six of Advent of Code
func PartOne() string {
	disc := ParseDisc(input)
	return disc.Name
}

// PartTwo Solves the second part of day six of Advent of Code
func PartTwo() int {
	d := ParseDisc(input)
	return FindNewWeight(&d)
}
