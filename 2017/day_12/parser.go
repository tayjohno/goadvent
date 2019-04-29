package main

import (
	"strconv"
	"strings"
)

type pipe struct {
	id          int
	connections []int
}

func parse(input string) (pipes []*pipe) {
	ss := strings.Split(input, "\n")
	for _, s := range ss {
		p := &pipe{}
		stringParts := strings.Split(s, " <-> ")
		p.id, _ = strconv.Atoi(stringParts[0])
		for _, intString := range strings.Split(stringParts[0], ", ") {
			conn, _ := strconv.Atoi(intString)
			p.connections = append(p.connections, conn)
		}
		pipes = append(pipes, p)
	}
	return
}
