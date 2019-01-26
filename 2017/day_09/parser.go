package main

import (
	"bytes"
	"fmt"
)

type node struct {
	garbage  bool // group or garbage
	contents string
	children []*node
}

func parse(i string) *node {
	buf := bytes.NewBufferString(i)
	r, _, _ := buf.ReadRune()
	if r != '{' {
		return &node{garbage: false}
	}
	return parseGroup(buf)
}

func parseGroup(buf *bytes.Buffer) *node {
	n := &node{garbage: false}
	for {
		r, _, err := buf.ReadRune()
		if err != nil {
			fmt.Println(err)
			return n
		}
		switch r {
		case '!':
			buf.ReadRune()
		case '{':
			n.addChild(parseGroup(buf))
		case '}':
			return n
		case '<':
			n.addChild(parseGarbage(buf))
		}
	}
}

func parseGarbage(buf *bytes.Buffer) (n *node) {
	n = &node{garbage: true}
	for {
		r, _, err := buf.ReadRune()
		if err != nil {
			fmt.Println(err)
			return
		}
		switch r {
		case '>':
			return
		case '!':
			buf.ReadRune()
		default:
			n.contents += string(r)
		}
	}
}

func (n *node) addChild(c *node) {
	i := len(n.children)
	if len(n.children) == cap(n.children) { // copy if needed
		old := n.children
		n.children = make([]*node, len(n.children), (cap(n.children)+1)*2)
		for i := range old {
			n.children[i] = old[i]
		}
	}
	n.children = n.children[:i+1]
	n.children[i] = c
}

func simplifiedString(n *node) (s string) {
	if n == nil {
		s += "nil"
	} else if n.garbage {
		s += "<>"
	} else {
		s += "{"
		for i := range n.children {
			s += simplifiedString(n.children[i])
		}
		s += "}"
	}
	return
}

func score(n *node) int {
	return nestedScore(n, 1)
}

func nestedScore(n *node, d int) int {
	s := d
	if n.garbage {
		return 0
	}
	for i := range n.children {
		s += nestedScore(n.children[i], d+1)
	}
	return s
}

func countGarbage(n *node) (c int) {
	if n.garbage {
		return len(n.contents)
	}
	for i := range n.children {
		c += countGarbage(n.children[i])
	}
	return c
}
