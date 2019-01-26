package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}

type registers map[string]int
type program []line

type line struct {
	com command
	pre predicate
}

type command struct {
	reg string
	ins string
	amt int
}

type predicate struct {
	reg string
	cmp string
	val int
}

// PartOne Solves the first part of day eight of Advent of Code
func PartOne() int {
	r := make(registers, 0)
	r.execute(input())
	return r.largest()
}

// PartTwo Solves the second part of day eight of Advent of Code
func PartTwo() int {
	r := make(registers, 0)
	return r.execute(input())
}

// Executes a line of code, running the command if the predicate passes
func (r *registers) execute(p program) int {
	if len(p) == 0 {
		return r.largest()
	}
	if r.checkPredicate(p[0].pre) {
		r.runCommand(p[0].com)
	}
	l1 := r.largest()
	l2 := r.execute(p[1:])
	if l1 > l2 {
		return l1
	}
	return l2
}

// Returns true if a predicate passes
func (r *registers) checkPredicate(p predicate) bool {
	switch p.cmp {
	case ">":
		return (*r)[p.reg] > p.val
	case ">=":
		return (*r)[p.reg] >= p.val
	case "==":
		return (*r)[p.reg] == p.val
	case "<=":
		return (*r)[p.reg] <= p.val
	case "<":
		return (*r)[p.reg] < p.val
	case "!=":
		return (*r)[p.reg] != p.val
	}
	fmt.Println("Oops!")
	return false
}

// Runs a single command, ignoring the predicate
func (r *registers) runCommand(c command) {
	switch c.ins {
	case "dec":
		(*r)[c.reg] -= c.amt
	default:
		(*r)[c.reg] += c.amt
	}
}

// Finds the largest value in any register
func (r *registers) largest() int {
	max := math.MinInt64
	for i := range *r {
		if (*r)[i] > max {
			max = (*r)[i]
		}
	}
	return max
}
