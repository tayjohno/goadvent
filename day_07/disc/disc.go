package disc

import (
	"adventofcode/day_07/input"
	"fmt"
	"strconv"
	"strings"
)

// ParseDisc returns the problem input for Day 7
func ParseDisc(input string) Disc {
	dict := defsToDiscDict(parseDefinitions(input))
	return buildTree(dict)
}

// definition is a disc that only references its children by name.
type definition struct {
	name       string
	weight     int
	childNames []string
}

// Disc is a type used for day 7 representing a Disc in memory.
type Disc struct {
	Name        string
	Weight      int
	childNames  []string
	Children    []*Disc
	parent      *Disc
	totalWeight int
}

type discDict map[string]*Disc

func defsToDiscDict(defs []definition) discDict {
	dict := make(discDict, len(defs))
	for i := range defs {
		dict[defs[i].name] = &Disc{
			Name:       defs[i].name,
			Weight:     defs[i].weight,
			childNames: defs[i].childNames,
			Children:   make([]*Disc, len(defs[i].childNames))}
	}
	return dict
}

func buildTree(dict discDict) Disc {
	for m := range dict {
		for i := range dict[m].childNames {
			n := dict[m].childNames[i]
			dict[m].Children[i] = dict[n]
			dict[n].parent = dict[m]
		}
	}
	for m := range dict {
		if dict[m].parent == nil {
			weigh(dict[m])
			return *dict[m]
		}
	}
	return Disc{}
}

// PrintDisc prints out the contents of the disc
func PrintDisc(disc Disc, depth int) {
	fmt.Print(strings.Repeat("-", depth))
	fmt.Printf("%s (%d) (%d)\n", disc.Name, disc.Weight, disc.totalWeight)
	for i := range disc.Children {
		PrintDisc(*disc.Children[i], depth+1)
	}
}

// FindNewWeight finds the imbalanced disc and returns the weight it should have. Needs at least 3
// branches to know which one is wrong, and what the delta is
func FindNewWeight(disc *Disc) int {
	weigh(disc)
	var goodWeight int
	if disc.Children[0].totalWeight == disc.Children[1].totalWeight || disc.Children[0].totalWeight == disc.Children[2].totalWeight {
		goodWeight = disc.Children[0].totalWeight
	} else {
		goodWeight = disc.Children[1].totalWeight
	}
	var badDisc *Disc
	var delta int
	for i := range disc.Children {
		if disc.Children[i].totalWeight != goodWeight {
			badDisc = disc
			delta = goodWeight - disc.Children[i].totalWeight
		}
	}
	return findNewWeightWithDelta(badDisc, delta)
}

func findNewWeightWithDelta(disc *Disc, delta int) int {
	if isBalanced(disc) {
		return disc.Weight + delta
	}
	// find disc off by delta
	for i := 1; i < len(disc.Children); i++ {
		if disc.Children[0].totalWeight-disc.Children[i].totalWeight == delta {
			return findNewWeightWithDelta(disc.Children[i], delta)
		}
	}
	return findNewWeightWithDelta(disc.Children[0], delta)
}

func isBalanced(disc *Disc) bool {
	for i := 1; i < len(disc.Children); i++ {
		if disc.Children[0].totalWeight != disc.Children[i].totalWeight {
			return false
		}
	}
	return true
}

func weigh(disc *Disc) int {
	if disc.totalWeight > 0 {
		return disc.totalWeight
	}
	disc.totalWeight = disc.Weight
	for i := range disc.Children {
		disc.totalWeight += weigh(disc.Children[i])
	}
	return disc.totalWeight
}

func parseDefinitions(s string) []definition {
	ss := strings.Split(input.Raw, "\n")
	d := make([]definition, len(ss))
	for i := range ss {
		d[i] = parseDefinition(ss[i])
	}
	return d
}

// parseDefinition returns a disc from a serialized string.
func parseDefinition(s string) definition {
	var n string
	var w int
	var c []string
	fields := strings.Fields(s)
	n = fields[0]
	w, _ = strconv.Atoi(strings.Trim(fields[1], "()"))
	if len(fields) > 3 {
		c = parseChildren(fields[3:])
	}
	return definition{n, w, c}
}

func printDefinition(d definition) {
	fmt.Printf("%s: weight=%d children=[%s]\n",
		d.name,
		d.weight,
		strings.Join(d.childNames, ", "))
}

func parseChildren(ss []string) []string {
	for i := range ss {
		ss[i] = strings.Trim(ss[i], ",")
	}
	return ss
}
