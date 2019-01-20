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
	Name       string
	Weight     int
	childNames []string
	Children   []*Disc
	parent     *Disc
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
			return *dict[m]
		}
	}
	return Disc{}
}

// PrintDisc prints out the contents of the disc
func PrintDisc(disc Disc, depth int) {
	fmt.Print(strings.Repeat("-", depth))
	fmt.Printf("%s (%d)\n", disc.Name, disc.Weight)
	for i := range disc.Children {
		PrintDisc(*disc.Children[i], depth+1)
	}
}

// ImbalancedDisc either returns a disc with an imbalance, or it returns the full weight.
func ImbalancedDisc(disc Disc) (*Disc, int) {
	weights := make([]int, len(disc.Children))
	totalWeight := 0
	for i := range disc.Children {
		d, w := ImbalancedDisc(*disc.Children[i])
		if d != nil {
			return d, 0
		}
		weights[i] = w
		totalWeight += w
	}
	for i := range weights {
		if weights[i] != totalWeight/len(weights) {
			return &disc, 0
		}
	}
	return nil, totalWeight + disc.Weight
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
