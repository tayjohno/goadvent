package main

import (
	"adventofcode/imath"
	"fmt"
	"strconv"
	"strings"
)

type banks [16]int
type banksLog map[banks]int

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}

// PartOne Solves the first part of day six of Advent of Code
func PartOne() int {
	i, _ := findLoop(startingBanks())
	return i
}

// PartTwo Solves the second part of day six of Advent of Code
func PartTwo() int {
	_, i := findLoop(startingBanks())
	return i
}

func startingBanks() banks {
	return deserialize(input)
}

func findLoop(b banks) (int, int) {
	log := banksLog{b: 0}
	i := 0
	for {
		i++
		b = cycle(b)
		if _, ok := log[b]; ok {
			break
		}
		log[b] = i
	}
	return len(log), i - log[b]
}

// Performs a single sycle on the banks
func cycle(b banks) banks {
	// Find Max
	max := 0
	maxI := 0
	for i := 0; i < 16; i++ {
		if b[i] > max {
			maxI = i
			max = b[i]
		}
	}
	// Distribute to Other Banks
	b[maxI] = 0
	i := maxI + 1
	for j := max; j > 0; j-- {
		b[imath.Mod(i, 16)]++
		i++
	}
	return b
}

func serialize(b banks) (s string) {
	for i := range b {
		s += strconv.Itoa(b[i])
		s += "\t"
	}
	s = strings.TrimSpace(s)
	return
}

func deserialize(s string) (b banks) {
	ss := strings.Split(s, "\t")
	for i := range ss {
		b[i], _ = strconv.Atoi(ss[i])
	}
	return
}

const input string = "14\t0\t15\t12\t11\t11\t3\t5\t1\t6\t8\t4\t9\t1\t8\t4"
