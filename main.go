package main

import (
	lem_in "lem_in/function"
)

func main() {
	ants, links, start, end := lem_in.ReadFile()
	graph := lem_in.ConvertToGraph(links)
	lem_in.Defs(end, graph, []string{start})
	lem_in.CheckErrors(ants, start, end, len(lem_in.Paths))
	lem_in.PrintAnts(ants, lem_in.Paths)
}
