package main

import (
	"fmt"
)

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}

// PartOne Solves the first part of day nine of Advent of Code
func PartOne() int {
	return score(parse(input))
}

// PartTwo Solves the second part of day nine of Advent of Code
func PartTwo() int {
	return countGarbage(parse(input))
}
