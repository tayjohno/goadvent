package main

import (
	"adventofcode/imath"
	"fmt"
)

func main() {
	fmt.Println(PartOne())
	fmt.Println(PartTwo())
}

// PartOne Solves the first part of day eleven of Advent of Code
func PartOne() int {
	directions := parseDirections(raw)
	return directions.distance()
}

// PartTwo Solves the second part of day eleven of Advent of Code
func PartTwo() (maxD int) {
	dirHash := make(directions)
	dirList := input(raw)
	for i := range dirList {
		dirHash[dirList[i]]++
		dirHash.simplify()
		tmpD := dirHash.distance()
		if tmpD > maxD {
			maxD = tmpD
		}
	}
	return maxD
}

func (d directions) simplify() {
	for {
		d.removeOpposites()
		d.mergeCorners()
		if d.simplified() {
			break
		}
	}
}

// A directions list is fully simplified if only 2 directions remain and they are side-by-side.
func (d directions) simplified() bool {
	remainingDirections := make([]direction, 0, 6)
	for i := n; i <= nw; i++ {
		if d[i] > 0 {
			l := len(remainingDirections)
			remainingDirections = remainingDirections[:l+1]
			remainingDirections[l] = i
		}
	}
	if len(remainingDirections) < 2 {
		return true
	} else if len(remainingDirections) > 2 {
		return false
	} else {
		diff := imath.Abs(int(remainingDirections[0] - remainingDirections[1]))
		if diff == 1 || diff == 5 {
			return true
		}
	}
	return false
}

func (d directions) removeOpposites() {
	for i := n; i <= se; i++ {
		min := imath.Min(d[n], d[s])
		d[n] -= min
		d[s] -= min
	}
}

func (d directions) mergeCorners() {
	for i := n; i <= nw; i++ {
		if d[i] > 0 && d[(i+2)%6] > 0 {
			min := imath.Min(d[i], d[(i+2)%6])
			d[i] -= min
			d[(i+1)%6] += min
			d[(i+2)%6] -= min
		}
	}
}

func (d directions) distance() (sum int) {
	d.simplify()
	for i := range allDirections {
		sum += d[allDirections[i]]
	}
	return
}

func (d directions) print() {
	fmt.Printf("n:%-6d", d[n])
	fmt.Printf("ne:%-6d", d[ne])
	fmt.Printf("se:%-6d", d[se])
	fmt.Printf("s:%-6d", d[s])
	fmt.Printf("sw:%-6d", d[sw])
	fmt.Printf("nw:%d\n", d[nw])
}
